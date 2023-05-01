import {LoginModel} from "$lib/models/classes/login.model";
import {TokenModel} from "$lib/models/classes/token.model";
import {sendCreateOrUpdate, sendGet} from "$lib/helpers/sender.helper";
import {authEP, countryEP, userEP} from "$lib/models/enums/endpoints.enum";
import {UserModel} from "$lib/models/classes/user.model";
import {currentUserStore, userStore} from "$lib/stores/user.store";
import {CountryModel} from "$lib/models/classes/country.model";
import {partCountryStore} from "$lib/stores/partCountry.store";

export async function loginAndGetUsers(payload:LoginModel) {
    await sendCreateOrUpdate<LoginModel, TokenModel>(authEP.LOGIN, payload, "POST").then(data => {
        if (data.body.token === "") {
            alert("Something went very wrong. Please refresh the page")
        }
    })

    await sendGet<Map<string,UserModel>>(userEP.ALL).then( userdata => {
        userStore.set(userdata.body)
        currentUserStore.set(userdata.body[payload.userId])
    })

    await sendGet<Array<CountryModel>>(countryEP.PARTICIPATING).then( countryData => {
        partCountryStore.set(countryData.body)
    })
}