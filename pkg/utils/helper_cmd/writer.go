package helper_cmd

type appLogWriter struct {
	msgChannel chan string
}

func (w *appLogWriter) Write(p []byte) (n int, err error) {
	w.msgChannel <- string(p)
	return len(p), nil
}
