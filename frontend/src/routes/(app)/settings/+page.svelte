<script lang="ts">
    import {currentUser} from "$lib/stores/user.store";
    import {staticEP} from "$lib/models/enums/endpoints.enum";
    import { enhance } from '$app/forms';
    import type {ActionData} from './$types';

    export let form:ActionData
    let hideNameForm = true

    function updateform(form: ActionData) {
        if (form != null) {
            hideNameForm = form.hideNameForm
        }
    }

    $: updateform(form)
</script>

<h1>Settings</h1>

{#if hideNameForm }
    <p>{$currentUser.name}</p>
    <form method="POST" action="?/showNameForm" use:enhance>
        <button>Edit</button>
    </form>
{:else}
    <form method="POST" action="?/updateName" use:enhance>
        <input type="text" name="name" bind:value={$currentUser.name} />
        <button>Save</button>
    </form>
{/if}
<img src={staticEP.IMG + $currentUser.icon} alt={$currentUser.name + "'s avatar"} style="max-width: 100px">