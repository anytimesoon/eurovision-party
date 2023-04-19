// eslint-disable-next-line @typescript-eslint/no-unused-vars,@typescript-eslint/ban-ts-comment
// @ts-ignore
import {sendCreateOrUpdate, sendGet} from "$lib/helpers/sender.helper";
import {LoginModel} from "$lib/models/classes/login.model";
import {TokenModel} from "$lib/models/classes/token.model";
import {authEP, countryEP, userEP} from "$lib/models/enums/endpoints.enum";

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
    let resp : TokenModel
    let payload = new LoginModel();
    payload.token = params.t;
    payload.userId = params.u;

    await sendCreateOrUpdate<LoginModel, TokenModel>(authEP.LOGIN, payload, "POST").then(data => {
        resp = data
        if (resp.body.token !== "") {
            localStorage.setItem("me", payload.userId)
        } else {
            alert("Something went very wrong. Please refresh the page")
        }
    })

    return {
        token: resp.token
    }
}



// async function loginAndGetUsers(payload){
//     let resp;
//
//
//
//     if (resp.error != "") {
//         //TODO error handling
//         alert(resp.error)
//         return
//     }
//
//     await sendGet<Map<string,UserModel>>(userEP.ALL).then( userdata => {
//         $userStore = userdata.body
//     })
//
//     await sendGet<Array<CountryModel>>(countryEP.PARTICIPATING).then( countryData => {
//         $partCountryStore = countryData.body
//     })
//
//     localStorage.setItem("me", JSON.stringify($userStore[payload.userId]))
//
//     if (JSON.parse(localStorage.getItem("me")).authLvl === 1 ) {
//         await goto("/admin/countries")
//     } else {
//         await goto("/")
//     }
//
// }