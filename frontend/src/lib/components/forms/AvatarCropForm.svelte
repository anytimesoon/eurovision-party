<script lang="ts">
    import Cropper from "svelte-easy-crop";
    import {ImageCropArea} from "$lib/models/classes/imageCropArea";
    import type {CropArea} from "svelte-easy-crop/types";
    import ImagePic from "svelte-material-icons/Image.svelte";
    import ContentSave from "svelte-material-icons/ContentSave.svelte";
    import {currentUser} from "$lib/stores/user.store";
    import {staticEP, userEP} from "$lib/models/enums/endpoints.enum";
    import {formButtonState} from "$lib/models/enums/formButtonState.enum";
    import FormButton from "$lib/components/buttons/FormButton.svelte";
    import {errorStore} from "$lib/stores/error.store";
    import ImageLoader from "$lib/components/images/ImageLoader.svelte";
    import {sessionStore} from "$lib/stores/session.store";

    const authorizedExtensions = ['image/jpg', 'image/jpeg', 'image/png']
    let cropArea:ImageCropArea = $state(new ImageCropArea())
    let img:string = $state(staticEP.AVATAR_IMG + $currentUser.icon)
    let formState = $state(formButtonState.DISABLED)
    let imageFiles:FileList = $state()
    let imageFile:File = $state()
    let aspect = $state(1)
    let controller:AbortController
    let fileName = ""
    interface Props {
        closer: VoidFunction;
    }

    let { closer }: Props = $props();

    let updateCrop = (e:CustomEvent) => {
        let pix:CropArea = e.detail.pixels
        cropArea.x = pix.x
        cropArea.y = pix.y
        cropArea.width = pix.width
        cropArea.height = pix.height
    }


    $effect(() => {
        if(imageFiles) {
            imageFile = imageFiles[0]

            if (authorizedExtensions.includes(imageFile.type)) {
                formState = formButtonState.ENABLED

                let reader = new FileReader()
                reader.onload = e => {
                    img = e.target.result as string
                }
                reader.readAsDataURL(imageFile)
            } else {
                $errorStore = "Only jpeg and png files are allowed"
            }
        }
    });

    const send = async (e: Event) => {
        e.preventDefault()
        formState = formButtonState.SENDING

        const cropped = await cropImage(imageFile, cropArea)
        const croppedAndResized = await resizeImage(cropped, 400)
        controller = new AbortController()

        const signal = controller.signal
        fileName = $currentUser.id + ".png"

        let fd = new FormData()
        fd.append('file', croppedAndResized, fileName)
        fd.append('id', $currentUser.id)

        const resp = await fetch(userEP.UPDATE_IMAGE,
            {
                method: "PUT",
                body: fd,
                signal: signal,
                headers: {
                    "Authorization": $sessionStore
                }
            }
        )
        if (!resp.ok) {
            $errorStore = "Oops... something went wrong. Please try another file"
            cancelUpload()
            closer()
        }

        $currentUser.icon = `${$currentUser.id}.png?${Date.now()}`
        closer()
        formState = formButtonState.DISABLED
    }

    export async function cropImage(file:File, pixelCrop:CropArea, rotation = 0):Promise<Blob> {
        const reader = new FileReader();
        const image = new Image()
        const canvas = document.createElement('canvas')
        const ctx = canvas.getContext('2d')
        const dataURItoBlob = (dataURI: string) => {
            const bytes = dataURI.split(',')[0].indexOf('base64') >= 0 ?
                atob(dataURI.split(',')[1]) :
                unescape(dataURI.split(',')[1]);
            const mime = dataURI.split(',')[0].split(':')[1].split(';')[0];
            const max = bytes.length;
            const ia = new Uint8Array(max);
            for (var i = 0; i < max; i++) ia[i] = bytes.charCodeAt(i);
            return new Blob([ia], {type:mime});
        };

        const crop = () => {
            const maxSize = Math.max(image.width, image.height)
            const safeArea = 2 * ((maxSize / 2) * Math.sqrt(2))

            // set each dimensions to double largest dimension to allow for a safe area for the
            // image to rotate in without being clipped by canvas context
            canvas.width = safeArea
            canvas.height = safeArea

            // translate canvas context to a central location on image to allow rotating around the center.
            ctx.translate(safeArea / 2, safeArea / 2)
            ctx.rotate((rotation * Math.PI) / 180)
            ctx.translate(-safeArea / 2, -safeArea / 2)

            // draw rotated image and store data.
            ctx.drawImage(
                image,
                safeArea / 2 - image.width * 0.5,
                safeArea / 2 - image.height * 0.5
            )
            const data = ctx.getImageData(0, 0, safeArea, safeArea)

            // set canvas width to final desired crop size - this will clear existing context
            canvas.width = pixelCrop.width
            canvas.height = pixelCrop.height

            // paste generated rotate image with correct offsets for x,y crop values.
            ctx.putImageData(
                data,
                Math.round(0 - safeArea / 2 + image.width * 0.5 - pixelCrop.x),
                Math.round(0 - safeArea / 2 + image.height * 0.5 - pixelCrop.y)
            )

            let dataUrl = canvas.toDataURL('image/png');
            return dataURItoBlob(dataUrl);
        }

        return new Promise((ok, no) => {
            if (!file.type.match(/image.*/)) {
                no(new Error("Not an image"));
                return;
            }

            reader.onload = (readerEvent: any) => {
                image.onload = () => ok(crop());
                image.src = readerEvent.target.result;
            };
            reader.readAsDataURL(file);
        })
    }

    const resizeImage = async (file:Blob, maxSize:number):Promise<Blob> => {
        const reader = new FileReader();
        const image = new Image();
        const canvas = document.createElement('canvas');
        const dataURItoBlob = (dataURI: string) => {
            const bytes = dataURI.split(',')[0].indexOf('base64') >= 0 ?
                atob(dataURI.split(',')[1]) :
                unescape(dataURI.split(',')[1]);
            const mime = dataURI.split(',')[0].split(':')[1].split(';')[0];
            const max = bytes.length;
            const ia = new Uint8Array(max);
            for (var i = 0; i < max; i++) ia[i] = bytes.charCodeAt(i);
            return new Blob([ia], {type:mime});
        };
        const resize = () => {
            let width = image.width;
            let height = image.height;

            if (width > height) {
                if (width > maxSize) {
                    height *= maxSize / width;
                    width = maxSize;
                }
            } else {
                if (height > maxSize) {
                    width *= maxSize / height;
                    height = maxSize;
                }
            }

            canvas.width = width;
            canvas.height = height;
            canvas.getContext('2d').drawImage(image, 0, 0, width, height);
            let dataUrl = canvas.toDataURL('image/png');
            return dataURItoBlob(dataUrl);
        };

        return new Promise((ok, no) => {
            if (!file.type.match(/image.*/)) {
                no(new Error("Not an image"));
                return;
            }

            reader.onload = (readerEvent: any) => {
                image.onload = () => ok(resize());
                image.src = readerEvent.target.result;
            };
            reader.readAsDataURL(file);
        })
    };

    function cancelUpload() {
        imageFiles = null
        fileName = ''
        controller.abort()
    }
</script>

<form onsubmit={e => send(e)}>
    <div class="h-60 w-60 relative mx-auto">
        {#if formState === formButtonState.DISABLED}
            <div class="p-3 overflow-hidden">
                <ImageLoader customClasses="w-full" src={staticEP.AVATAR_IMG + $currentUser.icon} alt={$currentUser.name + "'s avatar"}/>
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
        <div class="flex justify-between">
            <label for="avatar" class="cursor-pointer py-2 px-3 rounded text-typography-main">
                <span class="flex"><ImagePic size="1.4em"/> Browse</span>
                <input id="avatar" name="img" class="hidden" type="file" accept={authorizedExtensions.join(',')} bind:files={imageFiles}>
            </label>
            <FormButton buttonState={formState}>
                <ContentSave size="1.4em" /> Save
            </FormButton>
        </div>
    </div>

</form>
