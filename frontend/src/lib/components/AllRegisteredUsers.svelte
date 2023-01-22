<script lang="ts">
    import { onMount } from "svelte";
    import {sendGet} from '$lib/helpers/sender.helper';
    import {NewUserModel} from "$lib/models/classes/user.model";
    import {registeredUserStore} from "$lib/stores/user.store";
    import {userEP} from "$lib/models/enums/endpoints.enum";
    import {authEP, userFeEP} from "$lib/models/enums/endpoints.enum.js";

    onMount( () => {
        sendGet<NewUserModel[]>(userEP.REGISTERED).then( data => $registeredUserStore = data.body);
    })

</script>
    <table>
        <thead>
            <tr>
                <td>
                    Name
                </td>
                <td>
                    Email
                </td>
                <td>
                    Login
                </td>
            </tr>
        </thead>
        <tbody>
            {#each $registeredUserStore as user}
                <tr>
                <td>
                    <a href="{userFeEP.SINGLE_USER}{user.slug}">{user.name}</a>
                </td>
                <td>
                    {user.email}
                </td>
                <td>
                    <a href="{authEP.FE_LOGIN}{user.token}/{user.id}">Copy link</a>
                </td>
                </tr>
            {/each}
        </tbody>
    </table>
