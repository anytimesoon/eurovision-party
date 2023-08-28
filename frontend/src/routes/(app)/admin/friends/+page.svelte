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

<h1 class="text-center">Manage your friends</h1>
<div class="p-3">

    <form method="POST" action="?/register" use:enhance >
        <div>
            <label for="new-user-name">name</label>
            <input id="new-user-name" type="text" name="name" />
        </div>

        <!--        email <input id="" type="text" name="email" />-->
        <input type="submit">
    </form>
</div>

<div class="p-3">
    <table class="max-w-full min-w-full">
        <thead>
        <tr>
            <td>
                Name
            </td>
<!--            <td>-->
<!--                Email-->
<!--            </td>-->
            <td class="w-[30%] text-right">
                Login link
            </td>
        </tr>
        </thead>
        <tbody>
        {#each users as user}
            <tr class="my-3">
                <td>
                    <a href="{userSvelteEP.FIND_ONE}{user.slug}">{user.name}</a>
                </td>
<!--                <td>-->
<!--                    {user.email}-->
<!--                </td>-->
                <td class="text-right">
                    <button on:click={navigator.clipboard.writeText(authEP.SVELTE_LOGIN + "/" + user.token + "/" + user.id)}>Copy link</button>
                </td>
            </tr>
        {/each}
        </tbody>
    </table>
</div>