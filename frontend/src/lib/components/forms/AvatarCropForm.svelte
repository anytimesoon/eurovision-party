<script lang="ts">
    import Cropper from "svelte-easy-crop";
    import {ImageCropArea} from "$lib/models/classes/imageCropArea";
    import type {CropArea} from "svelte-easy-crop/types";
    import {currentUser} from "$lib/stores/user.store";
    import { enhance } from '$app/forms';
    import {staticEP} from "$lib/models/enums/endpoints.enum";

    let img:string = staticEP.IMG + $currentUser.icon
    let imageFiles:FileList
    let imageFile:File
    export let cropArea:ImageCropArea = new ImageCropArea()
    export let avatarForm:HTMLFormElement
    export let closeModal:VoidFunction
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

        let reader = new FileReader()
        reader.onload = e => {
            img = e.target.result as string
        }
        reader.readAsDataURL(imageFile)
    }
</script>

<form method="POST" action="?/updateImg" use:enhance enctype="multipart/form-data">
    <input type="hidden" name="id" bind:value={$currentUser.id}>
    <input type="hidden" name="x" bind:value={cropArea.x}>
    <input type="hidden" name="y" bind:value={cropArea.y}>
    <input type="hidden" name="height" bind:value={cropArea.height}>
    <input type="hidden" name="width" bind:value={cropArea.width}>


    <div class="h-60 w-60 relative mx-auto py-3">
        <Cropper
                image={img}
                bind:zoom={cropArea.zoom}
                bind:aspect
                on:cropcomplete={updateCrop}
                restrictPosition={true}
        />
    </div>


    <div class="w-60 mx-auto py-3 flex justify-between">
        <label for="avatar" class="cursor-pointer py-2 px-3 rounded text-typography-main">
            <i class="fa-regular fa-image"></i> Browse
            <input id="avatar" name="img" class="hidden" type="file" accept="image/png, image/jpeg" bind:files={imageFiles}>
        </label>
        <button type="submit"><i class="fa-regular fa-floppy-disk"></i> Save</button>
    </div>
</form>
