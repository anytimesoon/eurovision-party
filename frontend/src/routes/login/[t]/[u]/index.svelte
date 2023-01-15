<script type="ts">
	import { onMount } from "svelte";
    import {auth, userEP} from "$lib/models/enums/endpoints.enum"
	import { LoginModel } from '$lib/models/classes/login.model';
    import {sendCreateOrUpdate, sendGet} from '$lib/helpers/sender.helper';
	import type { TokenModel } from '$lib/models/classes/token.model';
    import {ResponseModel} from "$lib/models/classes/response.model";
    import {UserModel} from "$lib/models/classes/user.model";
    import {userStore} from "$lib/stores/user.store";

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

        await sendCreateOrUpdate<LoginModel, TokenModel>(auth.LOGIN, payload, "POST").then(data => {
            resp = data
            if (resp !== null) {
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

        if ($userStore[localStorage.getItem("me")].authLvl === 1 ) {
            window.location.replace("/countries")
        } else {
            window.location.replace("/")
        }
    }
</script>