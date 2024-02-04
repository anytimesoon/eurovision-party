<script lang="ts">
    import Cropper from "svelte-easy-crop";
    import {ImageCropArea} from "$lib/models/classes/imageCropArea";
    import type {CropArea} from "svelte-easy-crop/types";
    import {currentUser} from "$lib/stores/user.store";
    import { enhance } from '$app/forms';
    import {staticSvelteEP} from "$lib/models/enums/endpoints.enum";
    import {formButtonState} from "$lib/models/enums/formButtonState.enum";
    import FormButton from "$lib/components/forms/FormButton.svelte";

    const authorizedExtensions = ['image/jpg', 'image/jpeg', 'image/png']
    let cropArea:ImageCropArea = new ImageCropArea()
    let img:string = staticSvelteEP.IMG + $currentUser.icon
    let formState = formButtonState.DISABLED
    let imageFiles:FileList
    let imageFile:File
    let aspect = 1

    let updateCrop = (e:CustomEvent) => {
        let pix:CropArea = e.detail.pixels
        cropArea.x = pix.x
        cropArea.y = pix.y
        cropArea.width = pix.width
        cropArea.height = pix.height
    }


    $: if(imageFiles) {
        imageFile = imageFiles[0]

        if (authorizedExtensions.includes(imageFile.type)) {
            formState = formButtonState.ENABLED
            document.getElementById("avatar-upload-errors").innerText = ""

            let reader = new FileReader()
            reader.onload = e => {
                img = e.target.result as string
            }
            reader.readAsDataURL(imageFile)
        } else {
            document.getElementById("avatar-upload-errors").innerText = "Only jpeg and png are allowed"
        }
    }
</script>

<form method="POST" action="?/updateImg" enctype="multipart/form-data" use:enhance={() => {
        formState = formButtonState.SENDING

        return async ({ update }) => {
            await update()
            formState = formButtonState.DISABLED
        };
    }}>
    <input type="hidden" name="id" bind:value={$currentUser.id}>
    <input type="hidden" name="x" bind:value={cropArea.x}>
    <input type="hidden" name="y" bind:value={cropArea.y}>
    <input type="hidden" name="height" bind:value={cropArea.height}>
    <input type="hidden" name="width" bind:value={cropArea.width}>


    <div class="h-60 w-60 relative mx-auto">
        {#if formState === formButtonState.DISABLED}
            <div class="p-3 overflow-hidden">
                <img src={staticSvelteEP.IMG + $currentUser.icon} alt="Avatar"/>
            </div>
        {:else}
            <Cropper
                    image={img}
                    bind:zoom={cropArea.zoom}
                    bind:aspect
                    on:cropcomplete={updateCrop}
                    restrictPosition={true}
            />
        {/if}
    </div>



    <div class="w-60 mx-auto py-3 ">
        <div id="avatar-upload-errors" class="text-center text-typography-grey"></div>
        <div class="flex justify-between">
            <label for="avatar" class="cursor-pointer py-2 px-3 rounded text-typography-main">
                <i class="fa-regular fa-image"></i> Browse
                <input id="avatar" name="img" class="hidden" type="file" accept={authorizedExtensions.join(',')} bind:files={imageFiles}>
            </label>
            <FormButton state={formState}>
                <i class="fa-regular fa-floppy-disk"></i> Save
            </FormButton>
        </div>
    </div>

</form>
