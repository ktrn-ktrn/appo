function showAllInterviewer(str){
    let xhr4 = new XMLHttpRequest();
    xhr4.open('GET', '/interviewer');
    xhr4.onreadystatechange = function() {
        if (xhr4.status == 200 && xhr4.readyState == xhr4.DONE) {
            let res = JSON.parse(xhr4.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            $$(str).clearAll();
            $$(str).parse(xhr4.response);
        }
    }       
    xhr4.send();
}

function showInterviewer(){
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let xhr4 = new XMLHttpRequest();
    xhr4.open('GET', '/assessment/' + selectedAssessmentId + '/interviewer');
    xhr4.onreadystatechange = function() {
        if (xhr4.status == 200 && xhr4.readyState == xhr4.DONE) {
            let res = JSON.parse(xhr4.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            $$('interviewerList').clearAll();
            $$('interviewerList').parse(xhr4.response);
        }
    }       
    xhr4.send();
}

function showInterviewerById(){
    let xhr2 = new XMLHttpRequest();
    let selectedInterviewerId = $$('InterviewerDictionary').getSelectedItem().ID
    xhr2.open('GET', '/interviewer/' + selectedInterviewerId);
    xhr2.onreadystatechange = function() {
        if (xhr2.status == 200 && xhr2.readyState == xhr2.DONE) {
            let res = JSON.parse(xhr2.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            
            $$('editInterviewerForm').parse(res.Data);
        }
    }
    xhr2.send();
}

function RemoveInterviewer(){
    let selectedInterviewerId = $$('interviewerList').getSelectedItem().ID
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let xhr = new XMLHttpRequest();
    xhr.open('DELETE', '/assessment/' + selectedAssessmentId + '/interviewer/' + selectedInterviewerId);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showInterviewer(selectedAssessmentId)
            
        }
    }       
    xhr.send();
}

function RemoveInterviewerFromD(str){
    let selectedInterviewerId = $$(str).getSelectedItem().ID
    let xhr = new XMLHttpRequest();
    xhr.open('DELETE', '/interviewer/' + selectedInterviewerId);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showAllInterviewer(str)
            showAllInterviewer("popupList")
            showInterviewer()
        }
    }       
    xhr.send();
}

function AddInterviewer(surname, name, patronymic, emailPep, phoneNumber, position){
    let xhr = new XMLHttpRequest();
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let newInterviewer = {Surname: surname,
        Name: name,
        Patronymic: patronymic,
        Email: emailPep, 
        PhoneNumber: phoneNumber, 
        Position: position}
    xhr.open('PUT', '/assessment/' + selectedAssessmentId + '/interviewer');
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showInterviewer(selectedAssessmentId);
        }
    }       
    xhr.send(JSON.stringify(newInterviewer));
}

function AddInterviewerToD(str, surname, name, patronymic, emailPep, phoneNumber, position){
    let xhr = new XMLHttpRequest();
    let newInterviewer = {Surname: surname,
        Name: name,
        Patronymic: patronymic,
        Email: emailPep, 
        PhoneNumber: phoneNumber, 
        Position: position}
    xhr.open('PUT', '/interviewer');
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showAllInterviewer(str)
            showAllInterviewer("popupList")
        }
    }       
    xhr.send(JSON.stringify(newInterviewer));
}

function UpdateInterviewer(surname, name, patronymic, email, phoneNumber, position){
    let xhr = new XMLHttpRequest();
    let selectedInterviewerId = $$('InterviewerDictionary').getSelectedItem().ID
    let newInterviewer = {Surname: surname,
        Name: name,
        Patronymic: patronymic,
        Email: email, 
        PhoneNumber: phoneNumber, 
        Position: position}
    xhr.open('POST', '/interviewer/' + selectedInterviewerId);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showAllInterviewer("InterviewerDictionary")
            showAllInterviewer("popupList")
            showInterviewer()
        }
    }       
    xhr.send(JSON.stringify(newInterviewer));
}

