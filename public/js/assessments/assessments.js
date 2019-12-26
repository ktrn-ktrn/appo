let assessmentTable = {
    view: "datatable", 
        id: "assessments", 
        scroll:true, 
        width: 350, 
        editable: true,
        select: true,
        columns:[
            { id:"Status", template:"#StatusName#", header:"Статус", width: 130},
            { id:"Date", template:"#Date#", header:["Дата ассессмента", {content:"textFilter"}], sort:"string", width: 200},
        ], 
        on: {
            onItemClick: function(){
                showCandidate();
                showInterviewer();

                $$("delete").enable();
                $$("addPep").enable();
                $$("addSob").enable();
                $$("edit").enable();
                $$("status").enable();
                $$("deleteSob").enable();
            }}
                
}

var btnAddAssessment = {
    view:"button", type:"icon", icon:"wxi-plus", width: 50, click: function () {
        addAssess.show();
    }
}

var addAssess = webix.ui({
    view: 'window',
    head: 'Добавить ассессмент',
    id: "dateAssessEdit",
    width: 400,
    modal: true,
    close:true,
    position: 'center',
    body:{view:"form", id:"form1", scroll:false,  
        elements:[
        { view:"datepicker",
        id: "datePic",
        label: 'Дата ассессмента',
        labelWidth: 150,
        stringResult: true,
        timepicker: true
    },
        { view:"button", value:"Сохранить", click: addAssessment},
        ]}
});

var btnEditAssessment = {
    view:"button", type:"icon", icon:"wxi-pencil", width: 50, click: editAssessFunc
}

var btnRemoveAssessment = {
    view:"button", type:"icon", icon:"wxi-trash", width: 50, click: deleteAssessFunc
}

var btnStatusAssessment = {
    view:"button", type:"icon", icon:"wxi-dots", width: 50, click: assessmentStatusFunc
}

var editAssess = webix.ui({
    view: 'window',
    head: 'Изменить ассессмент',
    id: "dateAssessEdit",
    width: 400,
    modal: true,
    close:true,
    position: 'center',
    body:{view:"form", id:"forma", scroll:false,  
        elements:[
        { view:"datepicker",
        id: "datePica",
        label: 'Дата ассессмента',
        labelWidth: 150,
        name: "Date",
        stringResult: true,
        timepicker: true
    },
        { view:"button", value:"Сохранить", click: editAssessment},
        ]}
});

var askYouRemoveAssess = webix.ui({
    view: 'window',
    head: 'Удалить ассессмент?',
    id: "askRemoveAssess",
    width: 300,
    modal: true,
    position: 'center',
    body:{view:"form", id:"forma", scroll:false,
        elements:[
        {cols:[
            { view:"button", value:"Удалить", click: function(){
                RemoveAssessment();
                $$('peopleList').clearAll();
                $$('interviewerList').clearAll();
                $$("delete").disable();
                $$("addPep").disable();
                $$("addSob").disable();
                $$("edit").disable();
                $$("status").disable();
                $$("deleteSob").disable();
                this.getTopParentView().hide();
            }},
            { view:"button", value:"Отмена", click: function(){
                this.getTopParentView().hide(); }
            },
        ]}]
    }
});

var assessmentStatus = webix.ui({
    view:"popup",
    id:"assessStatus",
    position:"top",
    body:{view:"form", id:"form2", scroll:false,
        elements:[
            {view: "label", label: "Выбрать статус ассессмента:"},
            {view:"list",
            id: "assessStatusList",
            template:"#ID#. #Status#",
            autoHeight: true,
            select:true,
            on: {"onItemDblClick": function(){
                var statusItem = $$('assessStatusList').getSelectedItem()
                setAssessmentStatus(statusItem.ID, statusItem.Status);
                this.getTopParentView().hide();
            }}}
        ]}
});