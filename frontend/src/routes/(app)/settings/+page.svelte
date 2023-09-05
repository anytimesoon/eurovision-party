<script lang="ts">
    import {currentUser} from "$lib/stores/user.store";
    import {staticEP} from "$lib/models/enums/endpoints.enum";
    import { enhance } from '$app/forms';
    import type {ActionData} from './$types';
    import Cropper from 'svelte-easy-crop';
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import AdminNav from "$lib/components/AdminNav.svelte";
    import Modal from "$lib/components/Modal.svelte";
    import {onMount} from "svelte";

    let image = staticEP.IMG + $currentUser.icon
    let fileInput
    let crop = { x: 0, y: 0 }
    let zoom = 1
    let aspect = 1
    let modal:HTMLElement

    onMount(() => {
        modal = document.getElementById("modal")
    })

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

    function openModal() {
        modal.style.display = "block";
    }
</script>

{#if $currentUser.authLvl === authLvl.ADMIN }
    <AdminNav />
{/if}

<Modal />

<div class="px-1">
    <h1 class="text-center">Personal Settings</h1>
    <div class="py-3">
        <div class="max-w-max mx-auto">
        {#if hideNameForm }



                <form class="inline-block" method="POST" action="?/showNameForm" use:enhance>
                    <span class="inline-block">{$currentUser.name} <button><i class="fa-regular fa-pen-to-square"></i></button></span>
                </form>


        {:else}
            <form method="POST" action="?/updateName" use:enhance>
                <input type="text" name="name" bind:value={$currentUser.name} />
                <input type="hidden" name="id" bind:value={$currentUser.id} />
                <button><i class="fa-solid fa-floppy-disk"></i></button>
            </form>
        {/if}
        </div>
    </div>

    {#if hideImgForm }
        <div class="py-3 max-w-[10rem] mx-auto relative">
            <img class="w-full" src={staticEP.IMG + $currentUser.icon} alt={$currentUser.name + "'s avatar"}>
            <form method="POST" action="?/showImgForm" use:enhance>
<!--                <button class="absolute top-5 right-5"><i class="fa-regular fa-pen-to-square"></i></button>-->

            </form>
            <label for="avatar" class="absolute top-5 right-5">
                <i class="fa-regular fa-pen-to-square"></i>
                <input id="avatar" class="hidden" type="file" accept=".jpg, .jpeg, .png" on:change={openModal} bind:this={fileInput}>
            </label>
        </div>

    {:else}
        <div class="relative w-full max-h-full h-96">
            <div class="h-60">
                <Cropper
                        {image}
                        bind:crop
                        bind:zoom
                        bind:aspect
                        on:cropcomplete={e => console.log(e.detail)}
                />
            </div>

            <button class="absolute -bottom-10">Save</button>
        </div>
    {/if}
</div>