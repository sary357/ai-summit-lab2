package app

import "testing"

func TestCheckSystemStatus(t *testing.T) {
	// TODO: Please define how to check the system status
	// The following codes are just an example.
	status := CheckSystemStatus()
	if status.Status != "OK" {
		t.Error("CheckSystemStatus() failed, expected OK, got ", status.Status)
	}
}

func TestSaveAll(t *testing.T) {
       status := SaveAndExec("This is my codes", "requirements") 
       if status != true {
		t.Error("SaveAll failed, expected True, got ", status)
       }
         
}

