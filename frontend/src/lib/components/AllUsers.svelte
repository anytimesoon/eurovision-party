<script lang="ts">
    import {userEP} from "$lib/models/enums/endpoints.enum"
    import { onMount } from "svelte";
    import {sendGet} from '$lib/helpers/sender.helper';
    import {UserModel} from "$lib/models/classes/user.model";
    import {userStore} from "$lib/stores/user.store";
    import {staticEP} from "$lib/models/enums/endpoints.enum.js";

    onMount( () => {
        sendGet<Map<string, UserModel>>(userEP.ALL).then( data => $userStore = data.body);
    })

    const users = Object.entries($userStore)
</script>

<ul id="full-list" style="text-decoration: none;">
        {#each users as [id, user]}
            <li>
                <img src="{staticEP.IMG + user.icon}" alt="avatar missing" width="50">
                {user.name}
            </li>
        {/each}
</ul>