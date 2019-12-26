webix.ready(function(){
    webix.ui({
        rows:[
        {
            view:"toolbar", cols:[
                {view: "label", label: "Ассессменты"},
                {view: "button", label: "Выйти", href:"LogInWin.html",  width: 100, click: function(){
                    window.location.href = "LogInWin.html";
                    //window.show(this.config.href);
                    }
                }
            ]
        },
        {view:"tabview",
        animate:true,
        cells:[ 
            {header:"Ассессмент Менеджер",
            body:{
                rows:[
                    {cols:[
                        {rows:[
                            {cols:[
                            { view:"label", label: "<span class='webix_icon wxi-calendar' title='Calendar'></span>Ассессменты"},
                            btnAddAssessment,
                            btnEditAssessment,
                            btnRemoveAssessment,
                            btnStatusAssessment]},
                            assessmentTable
                        ]},
                    {view: "resizer"},
                    {rows:[
                        {cols:[
                        {view:"label", label: "<span class='webix_icon wxi-user' title='User'></span>Участники"},
                        btnAddPep,
                        btnEditPep,
                        btnRemovePep,
                        btnStatusPep]},
                        peopleTable,
                        {cols: [
                        { view:"label", label: "<span class='webix_icon wxi-user' title='User'></span>Сотрудники"},
                        btnAddInterviewer,
                        btnRemoveInterviewer,
                    ]},
                        interviewerTable,
                    ]}
                    ]}  
                ]
            }
        },
        {
            header:"Сотрудники",
            body:{
                rows:[
                    {cols:[
                        {view:"label", label: "<span class='webix_icon wxi-user' title='User'></span>Сотрудники"},
                        btnAddIntToDictionary,
                        btnRemoveIntFromDictionary,
                        btnEditIntInDictionary
                    ]},
                    AllInterviewer
                ]
            }
          }
    ]}]
    })
    showAllInterviewer("InterviewerDictionary")
    showAssessment();
    //showAssessmentStatusByID();
});



