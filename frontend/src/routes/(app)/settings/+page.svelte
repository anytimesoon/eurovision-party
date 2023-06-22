<script lang="ts">
    import {currentUser} from "$lib/stores/user.store";
    import {staticEP} from "$lib/models/enums/endpoints.enum";
    import { enhance } from '$app/forms';
    import type {ActionData} from './$types';
    import Cropper from 'svelte-easy-crop';

    let image = staticEP.IMG + $currentUser.icon
    let crop = { x: 0, y: 0 }
    let zoom = 1
    let aspect = 1

    export let form:ActionData
    let hideNameForm = true
    let hideImgForm = true

    function updateform(form: ActionData) {
        if (form != null) {
            hideNameForm = form.hideNameForm
            hideImgForm = form.hideImgForm
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
        <input type="hidden" name="id" bind:value={$currentUser.id} />
        <input type="hidden" name="email" bind:value={$currentUser.email} />
        <button>Save</button>
    </form>
{/if}

{#if hideImgForm }
    <img src={staticEP.IMG + $currentUser.icon} alt={$currentUser.name + "'s avatar"} style="max-width: 100px">
    <form method="POST" action="?/showImgForm" use:enhance>
        <button>Edit</button>
    </form>
{:else}
    <Cropper
            {image}
            bind:crop
            bind:zoom
            bind:aspect
            on:cropcomplete={e => console.log(e.detail)}
    />
    <button type="button" on:click={async () => {image = await getCroppedImg(image, pixelCrop)}}>Crop!</button>
{/if}
