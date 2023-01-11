<script type="ts">
	import { onMount } from "svelte";
	import { auth } from "$lib/models/enums/endpoints.enum"
	import { LoginModel } from '$lib/models/classes/login.model';
	import { sendCreateOrUpdate } from '$lib/helpers/sender.helper';
	import type { TokenModel } from '$lib/models/classes/token.model';
    import {ResponseModel} from "$lib/models/classes/response.model";

	onMount( () => {
		let path = window.location.pathname;
		let params:string[] = path.split('/');

		let payload = new LoginModel();
		payload.token = params[2];
		payload.userId = params[3];

        let resp:ResponseModel<TokenModel>
		sendCreateOrUpdate<LoginModel, TokenModel>(auth.LOGIN, payload, "POST").then(data => {
            resp = data
            if (resp !== null) {
                if (resp.error !== "") {
                    alert(resp.error)
                } else {
                    window.location.replace("/countries")
                }
            } else {
                alert("Something went very wrong. Please refresh the page")
            }
        })
	});
</script>