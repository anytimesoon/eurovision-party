<style>
    .spinner {
        position: fixed;
        z-index: 1;
        left: 0;
        right: 0;
        top: 0;
        bottom: 0;
        background: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='40' height='40' viewBox='0 0 50 50'%3E%3Cpath d='M28.43 6.378C18.27 4.586 8.58 11.37 6.788 21.533c-1.791 10.161 4.994 19.851 15.155 21.643l.707-4.006C14.7 37.768 9.392 30.189 10.794 22.24c1.401-7.95 8.981-13.258 16.93-11.856l.707-4.006z'%3E%3CanimateTransform attributeType='xml' attributeName='transform' type='rotate' from='0 25 25' to='360 25 25' dur='0.6s' repeatCount='indefinite'/%3E%3C/path%3E%3C/svg%3E") center / 50px no-repeat;
    }
</style>

<script type="ts">
	import { onMount } from "svelte";
    import {authEP, countryEP, userEP} from "$lib/models/enums/endpoints.enum"
	import { LoginModel } from '$lib/models/classes/login.model';
    import {sendCreateOrUpdate, sendGet} from '$lib/helpers/sender.helper';
	import type { TokenModel } from '$lib/models/classes/token.model';
    import {ResponseModel} from "$lib/models/classes/response.model";
    import {UserModel} from "$lib/models/classes/user.model";
    import {userStore} from "$lib/stores/user.store";
    import {goto} from "$app/navigation";
    import {CountryModel} from "$lib/models/classes/country.model";
    import {partCountryStore} from "$lib/stores/partCountry.store";

	onMount( () => {
		let path = window.location.pathname;
		let params:string[] = path.split('/');

		let payload = new LoginModel();
		payload.token = params[2];
		payload.userId = params[3];

        loginAndGetUsers(payload)

	});

    async function loginAndGetUsers(payload:LoginModel){
        let resp : ResponseModel<TokenModel>;

        await sendCreateOrUpdate<LoginModel, TokenModel>(authEP.LOGIN, payload, "POST").then(data => {
            resp = data
            if (resp.body.token !== "") {
                localStorage.setItem("me", payload.userId)
            } else {
                alert("Something went very wrong. Please refresh the page")
            }
        })

        if (resp.error != "") {
            //TODO error handling
            alert(resp.error)
            return
        }

        await sendGet<Map<string,UserModel>>(userEP.ALL).then( userdata => {
            $userStore = userdata.body
        })

        await sendGet<Array<CountryModel>>(countryEP.PARTICIPATING).then( countryData => {
            $partCountryStore = countryData.body
        })

        localStorage.setItem("me", JSON.stringify($userStore[payload.userId]))

        if (JSON.parse(localStorage.getItem("me")).authLvl === 1 ) {
            await goto("/admin/countries")
        } else {
            await goto("/")
        }

    }


</script>

<div class="spinner"></div>