webix.ready(function(){   
    webix.ui({  
        view:"window",
        id:"loginWindow",
        scroll:false,
        head: 'Авторизация',
        position: 'center',
        width:400,
        body:{ view:"form", id:"loginForm", scroll:false,
            elements:form,
            rules:{
                $all:webix.rules.isNotEmpty,
                emaleName: function(value){ return value == "k@gmail.com"; },
                passName:  function(value){ return value == "password"; }
            },},
        
    }).show()
});

var form = [
    { view:"text", id: "loginEmail", labelWidth: 100, name: "emaleName", label:"Email"},
    { view:"text", id: "loginPassword", labelWidth: 100, name: "passName",type:"password", label:"Password"},
    { view:"button", label:"Login",  href:"assessApp.html", click: function(){
        var form1 = this.getParentView();
        if (form1.validate()){
           window.location.href = "assessApp.html";
        }
        else{
            webix.message({ type:"error", text: "Form data is invalid" });
        }
    }}
]