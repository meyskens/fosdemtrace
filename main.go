package main

import (
	"context"
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	orderID := "0"

	ctx := context.Background()
	ctx, task := trace.NewTask(ctx, "makeCappuccino")
	trace.Start(os.Stderr)
	defer trace.Stop()
	trace.Log(ctx, "orderID", orderID)

	milk := make(chan bool)
	espresso := make(chan bool)
	done := make(chan bool)
	go func() {
		trace.WithRegion(ctx, "steamMilk", steamMilk)
		milk <- true
	}()
	go func() {
		trace.WithRegion(ctx, "extractCoffee", extractCoffee)
		espresso <- true
	}()
	go func() {
		defer task.End() // When assemble is done, the order is complete.
		<-espresso
		<-milk
		trace.WithRegion(ctx, "mixMilkCoffee", mixMilkCoffee)
		done <- true
	}()

	<-done
}

func steamMilk() {
	fmt.Println("Warming water...")
	time.Sleep(time.Second)
	fmt.Println("Outputting steam...")
	time.Sleep(2 * time.Second)
	fmt.Println("Milk is hot")
}

func extractCoffee() {
	fmt.Println("Crushing beans..")
	time.Sleep(time.Second)
	fmt.Println("Putting water on them...")
	time.Sleep(2 * time.Second)
	fmt.Println("Made an espresso shot")
	time.Sleep(time.Second)
	fmt.Println("Crushing beans...")
	time.Sleep(time.Second)
	fmt.Println("Putting water on them...")
	time.Sleep(2 * time.Second)
	fmt.Println("Made an espresso shot")
}

func mixMilkCoffee() {
	fmt.Println("Pooring milk...")
	fmt.Println("Trying lattee art...")
	time.Sleep(time.Second)
	fmt.Println("Failed lattee art :(")
}
