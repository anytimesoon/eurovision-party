<script lang="ts">
    import Spinner from "$lib/components/Spinner.svelte";
    import {botId, currentUser} from "$lib/stores/user.store";
    import {sessionStore} from "$lib/stores/session.store";
    import type { PageData } from './$types';
    import {goto} from "$app/navigation";
    import {onMount} from "svelte";
    import {authEP} from "$lib/models/enums/endpoints.enum";
    import {LoginModel} from "$lib/models/classes/login.model";
    import type {SessionModel} from "$lib/models/classes/session.model";
    import {UserModel} from "$lib/models/classes/user.model";
    import {post} from "$lib/utils/genericFetch";

    interface Props {
        data: PageData;
    }

    let { data }: Props = $props();

    onMount(async () => {
        const result = await post<SessionModel, LoginModel>(authEP.LOGIN, new LoginModel(data.loginToken, data.userId))

        $currentUser = UserModel.deserialize(result.user)
        $botId = result.botId
        $sessionStore = result.token

        if ($currentUser.isAdmin()) {
            await goto("/settings/countries")
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
