package main

import (
	"log"
	"time"

	sdk "github.com/gaia-pipeline/gosdk"
)

func CreateUser(args sdk.Arguments) error {
	log.Println("CreateUser has been started!")
	out, e, err := sh("pwd")
	if err != nil {
		return err
	}
	if e != nil {
		log.Printf("exit-error: %v\n", e)
		return e
	}
	log.Println(out)


	out, e, err = sh("hostname")
	if err != nil {
		return err
	}
	if e != nil {
		log.Printf("exit-error: %v\n", e)
		return e
	}
	log.Println(out)


	log.Println("os version")
	out, e, err = sh("cat /etc/os-release")
	if err != nil {
		return err
	}
	if e != nil {
		log.Printf("exit-error: %v\n", e)
		return e
	}
	log.Println(out)


	log.Println("xyz")
	out, e, err = sh("cat /data/repo/xyz")
	if err != nil {
		return err
	}
	if e != nil {
		log.Printf("exit-error: %v\n", e)
		return e
	}
	log.Println(out)



	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateUser has been finished!")
	return nil
}

func MigrateDB(args sdk.Arguments) error {
	log.Println("MigrateDB has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("MigrateDB has been finished!")
	return nil
}

func CreateNamespace(args sdk.Arguments) error {
	log.Println("CreateNamespace has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateNamespace has been finished!")
	return nil
}

func CreateDeployment(args sdk.Arguments) error {
	log.Println("CreateDeployment has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateDeployment has been finished!")
	return nil
}

func CreateService(args sdk.Arguments) error {
	log.Println("CreateService has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateService has been finished!")
	return nil
}

func CreateIngress(args sdk.Arguments) error {
	log.Println("CreateIngress has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateIngress has been finished!")
	return nil
}

func Cleanup(args sdk.Arguments) error {
	log.Println("Cleanup has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("Cleanup has been finished!")
	return nil
}

func main() {
	// Serve
	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}

	// CreateUser(nil)
}
