function showCandidate(){
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let xhr3 = new XMLHttpRequest();
    xhr3.open('GET', '/assessment/' + selectedAssessmentId + '/candidate/');
    xhr3.onreadystatechange = function() {
        if (xhr3.status == 200 && xhr3.readyState == xhr3.DONE) {
            let res = JSON.parse(xhr3.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            $$('peopleList').clearAll();
            $$('peopleList').parse(xhr3.response);
        }
    } 
    xhr3.send();
}

function showCandidateStatus(){
    let xhr = new XMLHttpRequest();
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let selectedCandidateId = $$('peopleList').getSelectedItem().ID
    xhr.open('GET', '/assessment/' + selectedAssessmentId + '/candidate/' + selectedCandidateId + '/status/');
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            if(res.Data == ''){
                $$("btnStatusPep").disable();
            } else{
            $$('peopleStatusList').clearAll();
            $$('peopleStatusList').parse(res.Data)
        }
        }
    }
    xhr.send();
}

function showCandidateById(){
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let selectedCandidateId = $$('peopleList').getSelectedItem().ID
    let xhr2 = new XMLHttpRequest();
    xhr2.open('GET', '/assessment/' + selectedAssessmentId + '/candidate/' + selectedCandidateId);
    xhr2.onreadystatechange = function() {
        if (xhr2.status == 200 && xhr2.readyState == xhr2.DONE) {
            let res = JSON.parse(xhr2.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            console.log("КАНДИДАТ: ", res.Data); 
            
            $$('editForm').parse(res.Data);
            $$("rebirthDatePep").setValue(new Date(res.Data.BirthDate));
        }
    }
    xhr2.send();
}

function RemoveCandidate(){
    let selectedCandidateId = $$('peopleList').getSelectedItem().ID
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let xhr = new XMLHttpRequest();
    xhr.open('DELETE', '/assessment/' + selectedAssessmentId + '/candidate/' + selectedCandidateId);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            $$("peopleList").remove(selectedCandidateId);
            showCandidate(selectedAssessmentId)
        }
    }       
    xhr.send();
}
/*
function AddCandidate(surname, name, patronymic, emailPep, phoneNumber, resume, addres, education, birthDate){
    let xhr = new XMLHttpRequest();
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let newCandidate = {Surname: surname,
        Name: name,
        Patronymic: patronymic,
        Email: emailPep, 
        PhoneNumber: phoneNumber, 
        Resume: resume,
        Addres: addres,
        Education: education,
        BirthDate: birthDate}
    xhr.open('PUT', '/assessment/' + selectedAssessmentId + '/candidate/');
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*
            showCandidate(selectedAssessmentId);
            console.log("НОВЫЙ КАНДИДАТ CALL: ", newCandidate)
        }
    }       
    xhr.send(JSON.stringify(newCandidate));
}*/

function AddCandidate(surname, name, patronymic, emailPep, phoneNumber, resume, addres, birthDate, education){
    let xhr = new XMLHttpRequest();
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let newCandidate = {Surname: surname,
        Name: name,
        Patronymic: patronymic,
        Email: emailPep, 
        PhoneNumber: phoneNumber, 
        Resume: resume,
        Addres: addres,
        BirthDate: birthDate,
        Education: education}
    xhr.open('PUT', '/assessment/' + selectedAssessmentId + '/candidate/');
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showCandidate(selectedAssessmentId);
            console.log("НОВЫЙ КАНДИДАТ CALL: ", newCandidate)
        }
    }       
    xhr.send(JSON.stringify(newCandidate));
}

function UpdateCandidate(surname, name, patronymic, emailPep, phoneNumber, resume, addres, birthDate, education){
    let xhr = new XMLHttpRequest();
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let selectedCandidateId = $$('peopleList').getSelectedItem().ID
    let newCandidate = {Surname: surname,
        Name: name,
        Patronymic: patronymic,
        Email: emailPep, 
        PhoneNumber: phoneNumber, 
        Resume: resume,
        Addres: addres,
        BirthDate: birthDate,
        Education: education}
    xhr.open('POST', '/assessment/' + selectedAssessmentId + '/candidate/' + selectedCandidateId);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showCandidate(selectedAssessmentId);
            console.log("ОБНОВЛЁННЫЙ КАНДИДАТ CALL: ", newCandidate)
        }
    }       
    xhr.send(JSON.stringify(newCandidate));
}

function setCandidateStatus(selectedStatusId, status){
    let xhr = new XMLHttpRequest();
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let selectedCandidateId = $$('peopleList').getSelectedItem().ID
    let newStatus = {//ID: idStatus,
        Status: status}
    xhr.open('POST', '/assessment/' + selectedAssessmentId + '/candidate/' + selectedCandidateId + '/status/' + selectedStatusId);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showCandidate(selectedAssessmentId);
        }
    }       
    xhr.send(JSON.stringify(newStatus));
}

function showAllCandidate(){
    let xhr4 = new XMLHttpRequest();
    xhr4.open('GET', '/candidate');
    xhr4.onreadystatechange = function() {
        if (xhr4.status == 200 && xhr4.readyState == xhr4.DONE) {
            let res = JSON.parse(xhr4.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            //$$('interviewerList').clearAll();
            $$('CandidateList').clearAll();
            $$('CandidateList').parse(xhr4.response);
        }
    }       
    xhr4.send();
}
