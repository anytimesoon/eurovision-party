<script lang="ts">
    import {currentUser, userStore} from "$lib/stores/user.store";
    import {staticSvelteEP} from "$lib/models/enums/endpoints.enum";
    import { enhance } from '$app/forms';
    import type {ActionData} from './$types';
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import AdminNav from "$lib/components/AdminNav.svelte";
    import Modal from "$lib/components/Modal.svelte";
    import AvatarCropForm from "$lib/components/forms/AvatarCropForm.svelte";

    export let form:ActionData
    let hideNameForm = true
    let openModal:VoidFunction
    let closeModal:VoidFunction
    let theme = localStorage.getItem("theme")
    let iconImage = staticSvelteEP.IMG + $currentUser.icon

    $: if(form){
        hideNameForm = form.hideNameForm

        if (form.user) {
            if(form.avatarUpdated){
                updateAvatar()
            }

            $currentUser = form.user
            closeModal()
        }
    }

    async function updateAvatar() {
        iconImage = staticSvelteEP.IMG + $currentUser.icon + `?${Date.now()}`
        $userStore[form.user.id].icon = iconImage
    }

    $: if(iconImage) {
        console.log(iconImage)
    }

    $: if(theme) {
        localStorage.setItem("theme", theme)
        document.querySelector("html")?.setAttribute("data-theme", theme)
    }
    function showNameForm() {
        hideNameForm = !hideNameForm
    }
</script>

{#if $currentUser.authLvl === authLvl.ADMIN }
    <AdminNav page="settings"/>
{/if}

<Modal bind:openModal={openModal} bind:closeModal={closeModal}>
    <AvatarCropForm />
</Modal>

<div class="pb-3">
    <h2 class="text-center">Personal Settings</h2>
    <div class="py-3">
        <div class="max-w-max mx-auto">
        {#if hideNameForm }
            <span class="inline-block text-2xl">{$currentUser.name} <button on:click={showNameForm} class="py-0 px-2"><i class="fa-regular fa-pen-to-square fa-2xs"></i></button></span>
        {:else}
            <form method="POST" action="?/updateName" use:enhance>
                <div class="w-fit mx-auto flex justify-center">
                    <input class="mr-3" type="text" name="name" bind:value={$currentUser.name} placeholder="Change your name"/>
                    <input type="hidden" name="id" bind:value={$currentUser.id} />
                    <button><i class="fa-solid fa-floppy-disk"></i></button>
                </div>
            </form>
        {/if}
        </div>
    </div>

    <div class="py-3 max-w-[10rem] mx-auto relative">
        <img class="w-full" src={iconImage} alt={$currentUser.name + "'s avatar"}>

        <button class="absolute top-5 right-2 cursor-pointer py-1 px-2 rounded" on:click={openModal}>
            <i class="fa-regular fa-pen-to-square"></i>
        </button>
    </div>

</div>

<div class="py-3">
    <h2 class="text-center">Themes</h2>
    <select bind:value={theme}>
        <option value="classic">Classic</option>
        <option value="light">Light</option>
    </select>
</div>
