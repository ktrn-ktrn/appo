var peopleTable = {
    view: "datatable", id: "peopleList", autoConfig: true, autoWidth: true, editable:true, select: true, columns:[
        { id:"ID", template:"#ID#", header:"#", width: 40},
        { id:"surname", template:"#Surname#", header:"Фамилия", width: 120},
        { id:"name", template:"#Name#", header:"Имя", width: 120},
        { id:"patronymic", template:"#Patronymic#", header:"Отчество", width: 120},
        { id:"email", template:"#Email#", header:"Email", width: 160},
        { id:"birthDate", template:"#BirthDate#", header: "Дата рождения", width: 200},
        { id:"phoneNumber",  template:"#PhoneNumber#", header:"Телефон", width: 130},
        { id:"status",  template:"#StatusName#", header:"Статус", width: 110}
    ], on: {onItemDblClick: editPeople}
}

var btnAddPep = {
    view:"button", id: "addPep", disabled:true, type:"icon", icon:"wxi-plus", width: 50, click: function () {
        askYouAddCandidate.show();
    }
}

var btnRemovePep = {
    view:"button", id: "delete", disabled:true, type:"icon", icon:"wxi-trash", width: 50, click: removeData
}

var btnEditPep = {
    view:"button", id: "edit", disabled:true, type:"icon", icon:"wxi-pencil", width: 50, click: editPeople
}

var btnStatusPep = {
    view:"button", id: "status", disabled:true, type:"icon", icon:"wxi-dots", width: 50, click: statusPeopleFunc
}

var addPep = webix.ui({
    view: 'window',
    head: 'Добавить участника',
    modal: true,
    width: 500,
    close:true,
    position: 'center',
    body: {
      view: 'form',
      id: "addForm",
      editable:true,
      elements: [
        { view: 'text', id:"surnamePep", label:"Фамилия"},
        { view: 'text', id:"namePep", label:"Имя"},
        { view: 'text', id:"patronymicPep", label:"Отчество"},
        { view: 'text', id: "emailPep", name: "email", label: 'Email'},
        { view: 'text', id: "phonePep", name: "phone", label: 'Телефон'},
        { view: 'text', id: "resumePep", name: "resume", label: 'Резюме'},
        { view: 'text', id: "addresPep", name: "addres", label: 'Адрес'},
        { view:"datepicker", id: "birthDatePep", name: "birthDate", stringResult: true, label: 'Дата рождения'},
        { view: 'text', id: "educationPep", name: "education", label: 'Образование'},
        {cols:[{ view: 'button', value: 'Добавить', click: addPeople},
        { view:"button", value:"Отмена", click:function(){
            $$("surnamePep").setValue("");
            $$("namePep").setValue("");
            $$("patronymicPep").setValue("");
            $$("emailPep").setValue("");
            $$("phonePep").setValue("");
            $$("resumePep").setValue("");
            $$("addresPep").setValue("");
            $$("birthDatePep").setValue("");
            $$("educationPep").setValue("");
            this.getTopParentView().hide(); 
          }}]}
      ]
    },
    move: true
});

var editPep = webix.ui({
    view: 'window',
    head: 'Изменить участника',
    id: "editPeople",
    modal: true,
    width: 500,
    editable: true,
    position: 'center',
    body: {
      view: 'form',
      id: "editForm",
      elements: [
        { view: 'text', id: "resurnamePep", name: "Surname", label: 'Фамилия'},
        { view: 'text', id: "renamePep",  name: "Name",label: 'Имя'},
        { view: 'text', id: "repatronymicPep",  name: "Patronymic",label: 'Отчество'},
        { view: 'text', id: "reemailPep",  name: "Email",label: 'Email'},
        { view: 'text', id: "rephonePep",  name: "PhoneNumber",label: 'Телефон'},
        { view: 'text', id: "reresumePep",  name: "Resume",label: 'Резюме'},
        { view: 'text', id: "readdresPep",  name: "Addres", label: 'Адрес'},
        { view: 'datepicker', id: "rebirthDatePep",  name: "BirthDate", stringResult: true, label: 'Дата рождения'},
        { view: 'text', id: "reeducationPep",  name: "Education",label: 'Образование'},
        {cols:[{ view: 'button', value: 'Изменить', click: editPep},
        { view:"button", value:"Отмена", click:function(){
            this.getTopParentView().hide(); 
          }}]}
      ]
    },
    move: true
});

var askYouRemoveCandidate = webix.ui({
    view: 'window',
    head: 'Удалить кандидата?',
    id: "askRemoveAssess",
    width: 300,
    modal: true,
    position: 'center',
    body:{view:"form", id:"forma", scroll:false,
        elements:[
        {cols:[
            { view:"button", value:"Удалить", click: function(){
                RemoveCandidate();
                this.getTopParentView().hide();
            }},
            { view:"button", value:"Отмена", click: function(){

                this.getTopParentView().hide(); }
            },
        ]}]
    }
});

var askYouAddCandidate = webix.ui({
    view: 'window',
    head: 'Добавить кандидата',
    id: "askRemoveAssess",
    width: 300,
    modal: true,
    close: true,
    position: 'center',
    body:{view:"form", id:"forma", scroll:false,
        elements:[
        {cols:[
            { view:"button", value:"Создать", click: function () {
                addPep.show();
                this.getTopParentView().hide();
            }},
            { view:"button", value:"Выбрать", click: function(){
                showAllCandidate();
                candidateList.show();
                this.getTopParentView().hide(); }
            },
        ]}]
    }
});

var peopleStatus = webix.ui({
    view:"popup",
    id:"pepStatus",
    autoHeight: true,
    position:"top",
    body:{view:"form", id:"form3", scroll:false,
        elements:[
            {view: "label", label: "Выбрать статус кандидата:"},
            {view:"list",
            id: "peopleStatusList",
            template:"#ID#. #Status#",
            select:true,
            on: {"onItemDblClick": function(){
                var statusItem = $$('peopleStatusList').getSelectedItem()
                setCandidateStatus(statusItem.ID, statusItem.Status);
                this.getTopParentView().hide();
            }}}
        ]}
});

var candidateList = webix.ui({
    view:"window",
    head: 'Выбрать кандидата',
    id:"addInterviewer",
    position: 'center',
    
    body:{view:"form", id:"formPep", width: 700, height:400,
        elements:[
            {view:"list",
            id: "CandidateList",
            template:"ФИО: #Surname# #Name# #Patronymic#",
            select:true,
        },
            {cols:[{ view: 'button', value: 'Добавить', click: function(){
                let selectedCandidateId = $$('CandidateList').getSelectedItem()
                console.log(selectedCandidateId);
                if(!$$("CandidateList").getSelectedId()){
                    webix.message("Кандидат не выбран из списка");
                    return;
                }
                else{
                    AddCandidate(selectedCandidateId.Surname, selectedCandidateId.Name, selectedCandidateId.Patronymic,
                    selectedCandidateId.Email, selectedCandidateId.PhoneNumber, selectedCandidateId.Resume, selectedCandidateId.Addres, 
                    selectedCandidateId.BirthDate, selectedCandidateId.Education);
                    this.getParentView().getParentView().hide()
                }
        }},
        { view:"button", value:"Отмена", click:function(){
            this.getTopParentView().hide(); 
          }},
        ]}
    ]}
});