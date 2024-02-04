<script lang="ts">
    import Spinner from "$lib/components/Spinner.svelte";
    import {botId, currentUser} from "$lib/stores/user.store";
    import type { PageData } from './$types';
    import {browser} from "$app/environment";
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import {goto} from "$app/navigation";
    import {loginURI} from "$lib/stores/loginURI.store";

    export let data: PageData;

    $: if(data) {
        if (browser) {
            $currentUser = data.currentUser
            $loginURI = data.loginToken + "/" + data.currentUser.id
            $botId = data.botId
            if ($currentUser.authLvl === authLvl.ADMIN && !data.hasLoggedIn) {
                goto("/admin/countries")
            } else {
                goto("/")
            }
        }
    }
</script>

<div class="h-screen flex flex-col justify-center">
    <div>
        <h3 class="text-center">Good evening Europe!</h3>
        <p class="text-center">Connecting you to this evening's entertainment</p>
        <Spinner />
    </div>

</div>
