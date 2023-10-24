<script lang="ts">
    import Spinner from "$lib/components/Spinner.svelte";
    import type { PageData } from './$types';
    import {onMount} from "svelte";
    import {authEP} from "$lib/models/enums/endpoints.enum";
    import type {ResponseModel} from "$lib/models/classes/response.model";
    import type {SessionModel} from "$lib/models/classes/session.model";
    import {goto} from "$app/navigation";
    import {currentUser} from "$lib/stores/user.store";
    import {authLvl} from "$lib/models/enums/authLvl.enum";

    export let data: PageData;

    onMount(async () => {
        let res = await fetch(`${authEP.INITIAL_DATA}/${data.token}/${data.userId}`)
        let session:ResponseModel<SessionModel> = await res.json()

        if(session.error != "") {
            await goto("/login")
            return
        }

        $currentUser = session.body.user

        if ($currentUser.authLvl === authLvl.ADMIN && !data.hasLoggedIn) {
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
