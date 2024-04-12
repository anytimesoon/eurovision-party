<script lang="ts">
    import {currentUser} from "$lib/stores/user.store";
    import {staticSvelteEP} from "$lib/models/enums/endpoints.enum";
    import { enhance } from '$app/forms';
    import type {ActionData} from './$types';
    import FileEditOutline from "svelte-material-icons/FileEditOutline.svelte";
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import AdminNav from "$lib/components/AdminNav.svelte";
    import Modal from "$lib/components/Modal.svelte";
    import AvatarCropForm from "$lib/components/forms/AvatarCropForm.svelte";
    import ContentSave from "svelte-material-icons/ContentSave.svelte";
    import FormButton from "$lib/components/forms/FormButton.svelte";
    import {formButtonState} from "$lib/models/enums/formButtonState.enum";
    import Toaster from "$lib/components/Toaster.svelte";

    export let form:ActionData
    let hideNameForm = true
    let openModal:VoidFunction
    let closeModal:VoidFunction
    let theme = localStorage.getItem("theme")
    let error:string = ""
    let closeToaster:VoidFunction
    let openToaster:VoidFunction

    $: if(form){
        hideNameForm = form.hideNameForm

        if (form.user) {
            form.user.icon += `?${Date.now()}`

            $currentUser = form.user
            closeModal()
        }
        if (form.error) {
            error = form.error
            openToaster()
        }
        form = null
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

<Toaster bind:openToaster={openToaster} bind:closeToaster={closeToaster}>
    {error}
</Toaster>

<Modal bind:openModal={openModal} bind:closeModal={closeModal}>
    <AvatarCropForm bind:error={error} bind:openToaster={openToaster}/>
</Modal>

<div class="pb-3">
    <h2 class="text-center">Personal Settings</h2>
    <div class="py-3">
        <div class="max-w-max mx-auto">
        {#if hideNameForm }
            <span class="inline-block text-2xl">{$currentUser.name}
                <button on:click={showNameForm} class="py-2 px-2">
                    <FileEditOutline size="0.75em"/>
                </button>
            </span>
        {:else}
            <form method="POST" action="?/updateName" use:enhance>
                <div class="w-fit mx-auto flex justify-center">
                    <input class="mr-3" type="text" name="name" bind:value={$currentUser.name} placeholder="Change your name"/>
                    <input type="hidden" name="id" bind:value={$currentUser.id} />
                    <FormButton state={formButtonState.ENABLED}>
                        <ContentSave size="1.4em" />
                    </FormButton>
                </div>
            </form>
        {/if}
        </div>
    </div>

    <div class="py-3 max-w-[10rem] mx-auto relative">
        <img class="w-full" src={staticSvelteEP.IMG + $currentUser.icon} alt={$currentUser.name + "'s avatr"}>

        <button class="absolute top-5 right-2 cursor-pointer py-2 px-2 rounded" on:click={openModal}>
            <FileEditOutline/>
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