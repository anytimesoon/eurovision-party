<script lang="ts">
    import {currentUser} from "$lib/stores/user.store";
    import {staticEP} from "$lib/models/enums/endpoints.enum";
    import { enhance } from '$app/forms';
    import type {ActionData} from './$types';
    import Cropper from 'svelte-easy-crop';
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import AdminNav from "$lib/components/AdminNav.svelte";

    let image = staticEP.IMG + $currentUser.icon
    let crop = { x: 0, y: 0 }
    let zoom = 1
    let aspect = 1

    export let form:ActionData
    let hideNameForm = true
    let hideImgForm = true

    function updateForm(form: ActionData) {
        if (form != null) {
            hideNameForm = form.hideNameForm
            hideImgForm = form.hideImgForm
        }
    }

    $: updateForm(form)
</script>

{#if $currentUser.authLvl === authLvl.ADMIN }
    <AdminNav />
{/if}

<div class="px-1">
    <h1>Personal Settings</h1>
    <div class="py-3">
        {#if hideNameForm }
            <div class="max-w-max mx-auto">

                <p class="inline-block">{$currentUser.name}</p>
                <form class="inline-block" method="POST" action="?/showNameForm" use:enhance>
                    <button>Edit</button>
                </form>
            </div>

        {:else}
            <form method="POST" action="?/updateName" use:enhance>
                <input type="text" name="name" bind:value={$currentUser.name} />
                <input type="hidden" name="id" bind:value={$currentUser.id} />
                <button>Save</button>
            </form>
        {/if}
    </div>

    {#if hideImgForm }
        <div class="py-3 max-w-max mx-auto">
            <img class="max-w-[10rem]" src={staticEP.IMG + $currentUser.icon} alt={$currentUser.name + "'s avatar"}>
            <form method="POST" action="?/showImgForm" use:enhance>
                <button>Edit</button>
            </form>
        </div>

    {:else}
        <div class="relative w-full max-h-full h-96">
            <Cropper
                    {image}
                    bind:crop
                    bind:zoom
                    bind:aspect
                    on:cropcomplete={e => console.log(e.detail)}
            />
            <button type="button" on:click={async () => {image = await getCroppedImg(image, pixelCrop)}}>Crop!</button>
        </div>
    {/if}
</div>