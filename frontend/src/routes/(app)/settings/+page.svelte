<script lang="ts">
    import {currentUser} from "$lib/stores/user.store";
    import {staticEP} from "$lib/models/enums/endpoints.enum";
    import { enhance } from '$app/forms';
    import type {ActionData} from './$types';
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import AdminNav from "$lib/components/AdminNav.svelte";
    import Modal from "$lib/components/Modal.svelte";
    import {ImageCropArea} from "$lib/models/classes/imageCropArea";

    export let form:ActionData
    let hideNameForm = true
    let openModal:VoidFunction
    let closeModal:VoidFunction
    let cropArea = new ImageCropArea()
    let imageFiles:FileList
    let imageFile:File
    let imageString:string
    let avatarForm:HTMLFormElement
    let theme:string

    $: if(form){
        hideNameForm = form.hideNameForm

        if (form.user) {
            $currentUser = form.user
            closeModal()
        }
    }

    $: if(imageFiles) {
        imageFile = imageFiles[0]

        let reader = new FileReader()
        reader.onload = e => {
            imageString = e.target.result as string
        }
        reader.readAsDataURL(imageFile)
        openModal()
    }

    $: if(theme) {
        document.querySelector("html")?.setAttribute("data-theme", theme)
    }
</script>

{#if $currentUser.authLvl === authLvl.ADMIN }
    <AdminNav page="settings"/>
{/if}

<Modal bind:openModal={openModal}
       bind:closeModal={closeModal}
       bind:cropArea={cropArea}
       img={imageString}
       avatarForm={avatarForm}/>

<div class="pb-3">
    <h2 class="text-center">Personal Settings</h2>
    <div class="py-3">
        <div class="max-w-max mx-auto">
        {#if hideNameForm }
                <form class="inline-block" method="POST" action="?/showNameForm" use:enhance>
                    <span class="inline-block text-2xl">{$currentUser.name} <button class="py-0 px-2"><i class="fa-regular fa-pen-to-square fa-2xs"></i></button></span>
                </form>
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
        <img class="w-full" src={staticEP.IMG + $currentUser.icon} alt={$currentUser.name + "'s avatar"}>

        <form method="POST" action="?/updateImg" use:enhance bind:this={avatarForm}>
            <input type="hidden" name="id" bind:value={$currentUser.id}>
            <input type="hidden" name="x" bind:value={cropArea.x}>
            <input type="hidden" name="y" bind:value={cropArea.y}>
            <input type="hidden" name="height" bind:value={cropArea.height}>
            <input type="hidden" name="width" bind:value={cropArea.width}>
<!--            <input type="hidden" name="zoom" bind:value={cropArea.zoom}>-->
            <label for="avatar" class="absolute top-5 right-2 cursor-pointer bg-primary py-1 px-2 rounded">
                <i class="fa-regular fa-pen-to-square"></i>
                <input id="avatar" name="img" class="hidden" type="file" accept="image/png, image/jpeg" bind:files={imageFiles}>
            </label>
        </form>
    </div>

</div>

<div class="py-3">
    <h2 class="text-center">Themes</h2>
    <select bind:value={theme}>
        <option value="classic">Classic</option>
        <option value="light">Light</option>
    </select>
</div>
