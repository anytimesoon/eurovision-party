<script lang="ts">
    import {currentUser} from "$lib/stores/user.store";
    import {staticEP} from "$lib/models/enums/endpoints.enum";
    import FileEditOutline from "svelte-material-icons/FileEditOutline.svelte";
    import AdminNav from "$lib/components/AdminNav.svelte";
    import Modal from "$lib/components/Modal.svelte";
    import AvatarCropForm from "$lib/components/forms/AvatarCropForm.svelte";
    import ImageLoader from "$lib/components/images/ImageLoader.svelte";
    import NameChangeForm from "$lib/components/forms/NameChangeForm.svelte";

    let hideNameForm = $state(true)
    let openModal:VoidFunction = $state()
    let closeModal:VoidFunction = $state()
    let theme = $state(localStorage.getItem("theme"))

    function formToggle() {
        hideNameForm = !hideNameForm
    }

    $effect(() => {
        if(theme) {
            localStorage.setItem("theme", theme)
            document.querySelector("html")?.setAttribute("data-theme", theme)
        }
    });


</script>

{#if $currentUser.isAdmin()}
    <AdminNav page="settings"/>
{/if}

<Modal bind:openModal={openModal} bind:closeModal={closeModal}>
    <AvatarCropForm closer={closeModal}/>
</Modal>

<div class="pb-3">
    <h2 class="text-center">Personal Settings</h2>
    <div class="py-3">
        <div class="max-w-max mx-auto">
        {#if hideNameForm}
            <span class="inline-block text-2xl">{$currentUser.name}
                <button onclick={formToggle} class="py-2 px-2">
                    <FileEditOutline size="0.75em"/>
                </button>
            </span>
        {:else}
            <NameChangeForm formToggle={formToggle} />
        {/if}
        </div>
    </div>

    <div class="w-[10rem] h-[10rem] mx-auto relative rounded overflow-hidden">
        <ImageLoader customClasses="w-full" src={staticEP.AVATAR_IMG + $currentUser.icon} alt={$currentUser.name + "'s avatar"}/>
        <button class="absolute top-2 right-2 cursor-pointer py-2 px-2 rounded" onclick={openModal}>
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