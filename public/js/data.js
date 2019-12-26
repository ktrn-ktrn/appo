idPepCounter = 5;
idAssessCounter = 3;
idInterviewerCounter = 3;

typesPep = [
    {idP:1, value: "Принят"},
    {idP:2, value: "Не принят"},
    {idP:3, value: "Записан"},
    {idP:4, value: "Пришел"},
    {idP:4, value: "Не пришел"}
];

typesAssess = [
    {idA:1, value: "Проведён"},
    {idA:2, value: "Не проведён"}
  ];

nameAssessment = "";

people = [
    {idPep: 1, name: "Имя Фамилия", email: "hohoho@gmail.com", phoneNumber: "89173276213", status: "--", resume: "Резюме 1", addres: "ул. Уличная 1, дом 31", birthDate: "18.08.1992", education: "СГУ, КНиИТ, компьютерная безопасность"},
    {idPep: 2, name: "Фамилия Имя", email: "grrg@gmail.com", phoneNumber: "89273441615", status: "--", resume: "Резюме 2", addres: "ул. Уличная 2, дом 32", birthDate: "28.07.1987", education: "СГУ, КНиИТ, педагогическое образование"},
    {idPep: 3, name: "Фамильная Фамилия", email: "sfadstf@gmail.com", phoneNumber: "89605513278", status: "--", resume: "Резюме 3", addres: "ул. Уличная 3, дом 33", birthDate: "30.12.1997", education: "СГУ, КНиИТ, компьютерная безопасность"},
    {idPep: 4, name: "Именное Имя", email: "dhdhdh@gmail.com", phoneNumber: "89878872387", status: "--", resume: "Резюме 4", addres: "ул. Уличная 4, дом 34", birthDate: "09.09.1990", education: "СГУ, КНиИТ, ФИиИТ"},
    {idPep: 5, name: "Артур", email: "arthur@gmail.com", phoneNumber: "88458819078", status: "--", resume: "Резюме 5", addres: "ул. Уличная 5, дом 35", birthDate: "14.04.1980", education: "СГУ, мех-мат, математика"},
]

assessments = [
    { idAssess: 1, date: "2019-07-26 16:00:00", status: "Не проведён", countPep: 5},
	{ idAssess: 2, date: "2019-05-07 15:30:00", status: "Не проведён", countPep: 5},
	{ idAssess: 3, date: "2019-04-13 16:00:00", status: "Не проведён", countPep: 5}
]

interviewer = [
    {idInterviewer: 1, name: "Собеседователь1"},
    {idInterviewer: 2, name: "Собеседователь2"},
    {idInterviewer: 3, name: "Собеседователь3"},
]