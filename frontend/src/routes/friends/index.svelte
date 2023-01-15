<script lang="ts">
    import {userEP} from "$lib/models/enums/endpoints.enum"
    import { onMount } from "svelte";
    import {sendGet} from '$lib/helpers/sender.helper';
    import {UserModel} from "$lib/models/classes/user.model";
    import {userStore} from "$lib/stores/user.store";
    import {updateUserStore} from "$lib/stores/user.store";
    import {staticEP} from "$lib/models/enums/endpoints.enum.js";

    onMount( () => {
        sendGet<UserModel[]>(userEP.ALL).then( data => updateUserStore(data.body));
    })

</script>

<ul id="full-list" style="text-decoration: none;">
    {#if [...$userStore].length > 0}
        {#each [...$userStore] as [id, user]}
                <li>
                    <img src="{staticEP.IMG + user.icon}" alt="{user.name}'s icon">
                    {user.name}
                </li>
        {/each}
    {/if}
</ul>