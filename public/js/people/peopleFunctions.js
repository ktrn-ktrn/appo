function addPeople(){
    if($$("namePep").getValue() != ""){
        console.log("НОВЫЙ КАНДИДАТ: ", $$("birthDatePep").getValue())
        AddCandidate($$("surnamePep").getValue(), $$("namePep").getValue(), $$("patronymicPep").getValue(), $$("emailPep").getValue(), 
        $$("phonePep").getValue(), $$("resumePep").getValue(), $$("addresPep").getValue(),  $$("birthDatePep").getValue(), 
        $$("educationPep").getValue());
        $$("surnamePep").setValue("");
        $$("namePep").setValue("");
        $$("patronymicPep").setValue("");
        $$("emailPep").setValue("");
        $$("phonePep").setValue("");
        $$("resumePep").setValue("");
        $$("addresPep").setValue("");
        $$("birthDatePep").setValue("");
        $$("educationPep").setValue("");
        this.getParentView().getParentView().getParentView().hide()
    }
    else{
        webix.message("Кандидат не может быть без имени");
    }
}

function removeData(){
    if(!$$("peopleList").getSelectedId()){
        webix.message("Кандидат не выбран");
        return;
    }

    if(!$$("assessments").getSelectedId()){
        webix.message("Ассессмент не выбран");
        return;
    }
    askYouRemoveCandidate.show();
}

function statusPeopleFunc(){
    if(!$$("peopleList").getSelectedId()){
        webix.message("Кандидат не выбран");
        return;
    }

    if(!$$("assessments").getSelectedId()){
        webix.message("Ассессмент не выбран");
        return;
    }

    showCandidateStatus()
    peopleStatus.show();
}

function editPeople(){
    if(!$$("peopleList").getSelectedId()){
        webix.message("Кандидат не выбран");
        return;
    }
    
    if(!$$("assessments").getSelectedId()){
        webix.message("Ассессмент не выбран");
        return;
    }
    
    showCandidateById()
    editPep.show();
}

function editPep(){
    if($$("renamePep").getValue() != ""){
        UpdateCandidate($$("resurnamePep").getValue(), 
        $$("renamePep").getValue(), 
        $$("repatronymicPep").getValue(), 
        $$("reemailPep").getValue(), 
        $$("rephonePep").getValue(), 
        $$("reresumePep").getValue(), 
        $$("readdresPep").getValue(), 
        $$("rebirthDatePep").getValue(), 
        $$("reeducationPep").getValue());
    }
    this.getParentView().getParentView().getParentView().hide()
}