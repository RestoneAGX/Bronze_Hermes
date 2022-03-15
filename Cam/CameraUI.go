package Cam

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
)

func OpenCam(origin *fyne.Window) int {
	CamOutput := canvas.Image{}
	CamOutput.FillMode = canvas.ImageFillOriginal
	CamOutput.Refresh()

	w := fyne.CurrentApp().NewWindow("Camera")
	w.SetContent(&CamOutput)
	w.Show()

	done := make(chan bool)
	defer close(done)

	var complete, evacuate bool
	w.SetOnClosed(func() {
		if !complete {
			evacuate = true
			done <- true
			return
		}
	})

	text := StartCamera(&CamOutput, done)
	complete = true

	if evacuate {
		fmt.Println("Evacuating...")
		return 0
	}

	w.Close()

	if text == "X" {
		dialog.ShowInformation("Time Up!", "The camera has been open for too long, but you can open it again.", *origin)
		return 0
	}

	conID, _ := strconv.Atoi(text)
	return conID
}