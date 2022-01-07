package helpers

import "flag"

func CommandLineArgs(wordlist_path, success_statuses, url, mode *string, threads *int) {
	flag.StringVar(wordlist_path, "wordlist", "/MUST/PROVIDE/STRING", "wordlist path")
	flag.StringVar(success_statuses, "success", "200", "success status (split by comma)")
	flag.StringVar(url, "url", "https://sans.org/", "URL to check")
	flag.StringVar(mode, "mode", "dir", "mode, [dir]")
	flag.IntVar(threads, "threads", 10, "number of threads to run")

	flag.Parse()
}
