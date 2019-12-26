function addAssessment(){
    if($$("datePic").getValue() != ""){
        AddAssessment($$("datePic").getValue())
        $$("datePic").setValue("");
        this.getParentView().getParentView().hide()
    }
}

function assessmentStatusFunc(){
    if(!$$("assessments").getSelectedId()){
        webix.message("Ассессмент не выбран");
        return;
    }
    showAssessmentStatus()
    assessmentStatus.show();
}

function editAssessFunc(){
    if(!$$("assessments").getSelectedId()){
        webix.message("Ассессмент не выбран");
        return;
    }
    showAssessmentById()
    editAssess.show();
}

function deleteAssessFunc(){
    if(!$$("assessments").getSelectedId()){
        webix.message("Ассессмент не выбран");
        return;
    }
    askYouRemoveAssess.show();
}

function editAssessment(){
    UpdateAssessment($$("datePica").getValue());
    this.getParentView().getParentView().hide()
}
