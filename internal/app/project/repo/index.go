package repo

import (
	"fmt"
	"lime/internal/app/project/model"
	"lime/internal/app/project/service"
	"lime/internal/global"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

func GetRepoInfo(info *model.ProjectInfo) error {
	// 仓库地址
	repoURL := info.RepoUrl
	// 临时克隆目录
	path := filepath.Join(global.ROOT_PATH, "temp")
	tmpDir, err := os.MkdirTemp(path, "git-repo")
	if err != nil {
		fmt.Printf("创建临时目录失败: %v\n", err)
		return err
	}
	defer os.RemoveAll(tmpDir)

	// 克隆仓库
	repo, err := git.PlainClone(tmpDir, false, &git.CloneOptions{
		URL:               repoURL,
		Progress:          os.Stdout,
		NoCheckout:        true,
		SingleBranch:      false,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		fmt.Printf("克隆仓库失败: %v\n", err)
		return err
	}

	remotes, err := repo.Remotes()
	if err != nil {
		fmt.Printf("获取远程仓库失败: %v\n", err)
		return err
	}

	fmt.Println("远程仓库信息:")
	for _, remote := range remotes {
		fmt.Printf("  %s\n", remote.Config().Name)

		urls := remote.Config().URLs
		for _, url := range urls {
			fmt.Printf("    %s\n", url)
		}

		// 获取远程仓库信息
		refs, err := remote.List(&git.ListOptions{})
		if err != nil {
			fmt.Printf("获取远程仓库信息失败: %v\n", err)
			return err
		}

		branchs := make([]model.BranchInfo, 0)
		tags := make([]model.TagInfo, 0)

		for _, ref := range refs {
			fmt.Printf("  %s\n", ref.Name())

			if strings.Contains(ref.Name().String(), "refs/tags") {
				tag := model.TagInfo{}
				tag.ProjectId = info.ID
				tag.Name = ref.Name().Short()
				tag.Hash = ref.Hash().String()
				tag.Description = ref.Name().String()
				tags = append(tags, tag)
			}

			// 获取分支信息
			if strings.Contains(ref.Name().String(), "refs/heads") {
				branch := model.BranchInfo{}
				branch.ProjectId = info.ID
				branch.Name = ref.Name().Short()
				branch.Hash = ref.Hash().String()
				branch.Description = ref.Name().String()
				branchs = append(branchs, branch)
			}
		}

		// 删除原有分支信息
		branchSrv := service.NewBranch()
		err = branchSrv.DeleteByFields(map[string]any{"project_id": info.ID})
		if err != nil {
			fmt.Printf("删除原有分支信息失败: %v\n", err)
			return err
		}

		// 删除原有 Tag 信息
		tagSrv := service.NewTag()
		err = tagSrv.DeleteByFields(map[string]any{"project_id": info.ID})
		if err != nil {
			fmt.Printf("删除原有 Tag 信息失败: %v\n", err)
			return err
		}

		// 保存分支信息
		err = branchSrv.CreateBatch(branchs)
		if err != nil {
			fmt.Printf("保存分支信息失败: %v\n", err)
			return err
		}

		// 保存 Tag 信息
		err = tagSrv.CreateBatch(tags)
		if err != nil {
			fmt.Printf("保存 Tag 信息失败: %v\n", err)
			return err
		}
	}

	return nil
}
