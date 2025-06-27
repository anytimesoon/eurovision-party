<script lang="ts">
    import Spinner from "$lib/components/Spinner.svelte";
    import {botId, currentUser} from "$lib/stores/user.store";
    import {sessionStore} from "$lib/stores/session.store";
    import type { PageData } from './$types';
    import {goto} from "$app/navigation";
    import {onMount} from "svelte";
    import {authEP} from "$lib/models/enums/endpoints.enum";
    import {LoginModel} from "$lib/models/classes/login.model";
    import type {ResponseModel} from "$lib/models/classes/response.model";
    import type {SessionModel} from "$lib/models/classes/session.model";
    import {redirect} from "@sveltejs/kit";
    import {UserModel} from "$lib/models/classes/user.model";

    interface Props {
        data: PageData;
    }

    let { data }: Props = $props();

    onMount(async () => {
        const res = await fetch(authEP.LOGIN, {
            method: "POST",
            body: JSON.stringify(new LoginModel(data.loginToken, data.userId))
        })

        const login: ResponseModel<SessionModel> = await res.json()
        if (login.error != "") {
            redirect(401, "/login")
        }

        $currentUser = UserModel.deserialize(login.body.user)
        $botId = login.body.botId
        $sessionStore = login.body.token

        if ($currentUser.isAdmin()) {
            await goto("/admin/countries")
        } else {
            await goto("/")
        }
    })
</script>

<div class="h-screen flex flex-col justify-center">
    <div>
        <h3 class="text-center">Good evening Europe!</h3>
        <p class="text-center">Connecting you to this evening's entertainment</p>
        <Spinner />
    </div>

</div>
