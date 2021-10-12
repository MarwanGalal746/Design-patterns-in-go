package main

import "Chain_of_responsipility/models"

func main() {
	var mgr models.Manager
	var vp models.VicePresident
	var p models.President
	mgr.SetSuccessor(&vp)
	vp.SetSuccessor(&p)
	p.SetSuccessor(nil)
	mgr.HandleRequest(models.Task)
	mgr.HandleRequest(models.Vacation)
}
