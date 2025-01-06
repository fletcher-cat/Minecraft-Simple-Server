package main

//This is cross-platform compatable. (Windows, Linux)

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/abdfnx/gosh"
)

func main() {
	system := runtime.GOOS //determine what OS is being used

	//deal with Java (unfinished)
	fmt.Println("Minecraft server setup\nGetting Java version")

	err, java, erroutput := gosh.RunOutput(`java --version`)

	if err != nil { //if error is not nothing, do this

		fmt.Println(erroutput)

		fmt.Println("Java is not installed or is not part of the system PATH. You can check if it is installed but not part of the system PATH by searching for 'Java', 'JRE', or 'JDK' in Programs and Features.\nOtherwise this program can download and install the latest version of Java 21 for you.\nThe latest Java 21 should be able to run any older version of Minecraft server as well.")

		fmt.Println("Install Java 21? Press 'y' to accept, any other key to decline.")

		var install_java string

		fmt.Scanln(&install_java)

		if system == "windows" && (install_java == "y" || install_java == "Y") { // || if system is running windows and the user agrees
			fmt.Println("Downloading Java. This may take a few minutes.")
			err, _, errout := gosh.RunOutput(`wget https://github.com/adoptium/temurin21-binaries/releases/download/jdk-21.0.5%2B11/OpenJDK21U-jre_x64_windows_hotspot_21.0.5_11.msi -o installer.msi`) //download the openjdk msi with powershell and name it installer.msi
			if err != nil {
				fmt.Println(errout)
			}

			var cont string
			fmt.Println("Downloaded. Please go into the folder this program is in and run 'installer.msi' to install Java 21. Accept all default settings.")
			fmt.Println("Once you have installed Java, type anything and press enter.")
			fmt.Scan(&cont)
			if cont != "" {
			}
		}
		if system == "linux" && (install_java == "y" || install_java == "Y") {
			fmt.Println("Because Linux comes in different distributions and installing Java would require root access, it must be manually installed, download Java 21 or the latest LTS.\nOnce that is done, you may resume the program.")
		}
	}

	fmt.Println(java)

	//Java works. Move on to server setup.
	var eula_accept string

	fmt.Println("Before proceeding, you must agree to the Minecraft EULA (https://www.minecraft.net/en-us/eula) and Microsoft Privacy Policy (https://www.microsoft.com/en-us/privacy/privacystatement).")

	fmt.Println("If you accept, press Y then enter. If you do not accept, type anything else to exit.")

	fmt.Scan(&eula_accept)

	if eula_accept != "Y" && eula_accept != "y" {
		fmt.Print("The terms must be accepted to proceed.\nExiting in 5 seconds.")
		time.Sleep(5 * time.Second) //wait 5 seconds then close
		os.Exit(0)
	}

	//user accepted eula, create the accepted eula file
	eula_file, err := os.Create("eula.txt")
	if err != nil {
		fmt.Println("There was an error creating the EULA text file. Do you have permission to modify or create files in this directory?")
	}
	eula_file.WriteString("eula=true")
	eula_file.Close()

	var jar_type int
	fmt.Println("You'll need a server .jar file in order to proceed. Download .jar automatically or use a different version? (Hit option number then enter.)")
	fmt.Println("(1) Automatically Download Vanilla Server Version 1.21.4\n(2) Use custom .jar, useful for newer/older version of Minecraft or running a Fabric server. Forge is not supported.")
	fmt.Scan(&jar_type)
	if jar_type != 1 && jar_type != 2 {
		fmt.Println("Please choose one of the options. (1/2)")
	}

	if jar_type == 1 && system == "windows" { //windows will not download the server jar properly without "-o server.jar"
		fmt.Println("Please wait for the .jar to finish downloading.")
		gosh.Run("wget https://piston-data.mojang.com/v1/objects/4707d00eb834b446575d89a61a11b5d548d8c001/server.jar -o server.jar")
		fmt.Println("Download successful.")
	}
	if jar_type == 1 && system == "linux" { //however using -o with linux causes a different issue, so omit it
		fmt.Println("Please wait for the .jar to finish downloading.")
		gosh.Run("wget https://piston-data.mojang.com/v1/objects/4707d00eb834b446575d89a61a11b5d548d8c001/server.jar")
		fmt.Println("Download successful.")
	}

	if jar_type == 2 {
		fmt.Println("Take the .jar file you want to use and rename it to 'server.jar'. Ensure it is not renamed to server.jar.jar, if you have hidden file extensions on simply rename it to 'server'. Place the renamed jar file in the same folder as this program and proceed.")
	}

	var ram_amount int

	fmt.Println("How much RAM would you like to assign to the server? Amount in GB, whole numbers only. Ensure you leave sufficient RAM for your operating system and other programs to run.")

	fmt.Scan(&ram_amount)

	if system == "windows" { //which OS the user is on changes the file extension

		bat_file, err := os.Create("start.bat") //create the bat

		if err != nil {

			fmt.Println("There was an error creating the batch file. Do you have permission to modify or create files in this directory?")

		}

		_, err = bat_file.WriteString(fmt.Sprintf("java -Xmx%dG -jar server.jar --nogui", ram_amount)) //write the ram amount the user gave into the script to start the server

		if err != nil {

			fmt.Println("There was an error writing to the batch file.")

		}
	}
	if system == "linux" {
		sh_file, err := os.Create("start.sh")

		if err != nil {

			fmt.Println("There was an error creating the script file. Do you have permission to modify or create files in this directory?")

		}

		_, err = sh_file.WriteString(fmt.Sprintf("java -Xmx%dG -jar server.jar --nogui", ram_amount))

		if err != nil {

			fmt.Println("There was an error writing to the script file.")

		}
		gosh.ShellCommand("chmod +x start.sh") //make the script executable
	}
	fmt.Println("Success! The file to start your server has been created. Please navigate to the folder this program is in and run start.bat on Windows or start.sh on Linux.")
	fmt.Println("You may now delete this program.")
}
