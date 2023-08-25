<script lang="ts">
    import { enhance } from '$app/forms';
    import {authEP, userSvelteEP} from "$lib/models/enums/endpoints.enum";
    import type {PageData, ActionData} from "./$types";
    import {currentUser} from "$lib/stores/user.store";
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import AdminNav from "$lib/components/AdminNav.svelte";

    export let data:PageData
    export let form:ActionData

    let users = data.users

    const updateUsers = (form) => {
        if(form != null) {
            users = [...users, form.user]
        }
    }

    $: updateUsers(form)
</script>

{#if $currentUser.authLvl === authLvl.ADMIN }
    <AdminNav />
{/if}

<form method="POST" action="?/register" use:enhance >
    name <input type="text" name="name" />
    email <input type="text" name="email" />
    <input type="submit">
</form>

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
    {#each users as user}
        <tr>
            <td>
                <a href="{userSvelteEP.FIND_ONE}{user.slug}">{user.name}</a>
            </td>
            <td>
                {user.email}
            </td>
            <td>
                <button on:click={navigator.clipboard.writeText(authEP.SVELTE_LOGIN + "/" + user.token + "/" + user.id)}>Copy link</button>
            </td>
        </tr>
    {/each}
    </tbody>
</table>