package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	deleteFileList = []string{
		"/Library/Java/JavaVirtualMachines",
		"/Library/PreferencePanes/JavaControlPanel.prefPane",
		"/Library/Internet Plug-Ins/JavaAppletPlugin.plugin",
		"/Library/LaunchAgents/com.oracle.java.Java-Updater.plist",
		"/Library/PrivilegedHelperTools/com.oracle.java.JavaUpdateHelper",
		"/Library/LaunchDaemons/com.oracle.java.JavaUpdateHelper.plist",
		"/Library/Preferences/com.oracle.java.Helper-Tool.plist",
		"/var/root/Library/Preferences/com.oracle.javadeployment.plist",
		"~/Library/Preferences/com.oracle.java.JavaAppletPlugin.plist",
		"~/Library/Preferences/com.oracle.javadeployment.plist",
		"~/.oracle_jre_usage",
		"/var/db/receipts/com.oracle.jre.bom",
		"/var/db/receipts/com.oracle.jre.plist",
	}
	deleteInDir = "/var/db/receipts"
	errors      []string
)

func main() {
	for _, v := range deleteFileList {
		err := os.RemoveAll(v)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}

	err := filepath.Walk(deleteInDir, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "com.oracle.jdk") {
			err := os.RemoveAll(path)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
		return nil
	})
	if err != nil {
		errors = append(errors, err.Error())
	}

	if len(errors) != 0 {
		fmt.Println("Error:")
		for _, v := range errors {
			fmt.Println(v)
		}
		fmt.Println("")
		fmt.Println("You wasn't been able to say \"bye bye\" to Java, sorry :(")
		os.Exit(1)
	}

	fmt.Println("Finished!")
	fmt.Println("You said \"bye bye\" to Java, congratulations :)")
}
