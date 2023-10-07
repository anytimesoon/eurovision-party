<script lang="ts">
    import Cropper from "svelte-easy-crop";
    import {ImageCropArea} from "$lib/models/classes/imageCropArea";
    import type {CropArea} from "svelte-easy-crop/types";

    export let img:string
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

    const sendImage = () => {
        avatarForm.requestSubmit()
    }
</script>

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
    <button on:click={sendImage}><i class="fa-regular fa-floppy-disk"></i> Save</button>
    <button on:click={closeModal} class="bg-transparent"><i class="fa-solid fa-ban"></i> Cancel</button>
</div>