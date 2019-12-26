function showAssessment(){
    let xhr = new XMLHttpRequest();
    xhr.open('GET', '/assessment');
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            $$('assessments').clearAll();
            $$('assessments').parse(res.Data)
        }
    }
    xhr.send();
}

function showAssessmentStatus(){
    let xhr = new XMLHttpRequest();
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    xhr.open('GET', '/assessment/' + selectedAssessmentId + '/status');
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            $$('assessStatusList').clearAll();
            $$('assessStatusList').parse(res.Data)
        }
    }
    xhr.send();
}

function showAssessmentById(){
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let xhr2 = new XMLHttpRequest();
    xhr2.open('GET', '/assessment/' + selectedAssessmentId);
    xhr2.onreadystatechange = function() {
        if (xhr2.status == 200 && xhr2.readyState == xhr2.DONE) {
            let res = JSON.parse(xhr2.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            $$("datePica").setValue(new Date(res.Data.Date));        
        }
    }
    xhr2.send();
}

function RemoveAssessment(){
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let xhr = new XMLHttpRequest();
    
    xhr.open('DELETE', '/assessment/' + selectedAssessmentId);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showAssessment();
        }
    }       
    xhr.send();
}

function AddAssessment(datePic){
    let xhr = new XMLHttpRequest();
    let newAssessment = {Date: datePic}
    xhr.open('PUT', '/assessment', true);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showAssessment();
        }
    }       
    xhr.send(JSON.stringify(newAssessment));
}

function UpdateAssessment(datePic){
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let xhr = new XMLHttpRequest();
    let newAssessment = {Date: datePic}
    xhr.open('POST', '/assessment/' + selectedAssessmentId, true);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showAssessment();
        }
    }       
    xhr.send(JSON.stringify(newAssessment));
}

function setAssessmentStatus(selectedStatusId, status){
    let selectedAssessmentId = $$('assessments').getSelectedItem().ID
    let xhr = new XMLHttpRequest();
    let newStatus = {//ID: idStatus,
    Status: status}
    xhr.open('POST', '/assessment/' + selectedAssessmentId + '/status/' + selectedStatusId, true);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showAssessment();
            //$$('assessments').select(selectedAssessmentId);
        }
    }       
    xhr.send(JSON.stringify(newStatus));
}