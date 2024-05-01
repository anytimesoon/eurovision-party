<script lang="ts">
    import {currentUser} from "$lib/stores/user.store";
    import Send from "svelte-material-icons/Send.svelte";
    import Image from "svelte-material-icons/Image.svelte";
    import {ChatMessageModel} from "$lib/models/classes/chatMessage.model";
    import {CommentModel} from "$lib/models/classes/comment.model";
    import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
    import {commentQueue} from "$lib/stores/commentQueue.store";
    import {replyComment} from "$lib/stores/replyComment.store";
    import ReplyComment from "$lib/components/chat/ReplyComment.svelte";
    import Dropzone from "dropzone";
    import {staticSvelteEP} from "$lib/models/enums/endpoints.enum";
    import {onMount} from "svelte";
    import CloseCircleOutline from "svelte-material-icons/CloseCircleOutline.svelte";
    import Spinner from "$lib/components/Spinner.svelte";

    let textArea:HTMLTextAreaElement
    let message:string = ""
    let gallery:HTMLDivElement
    let previewTemplate:HTMLDivElement
    let uploader:Dropzone
    let fileName:string = ""
    let isDisabled = false

    onMount(() => {
        const template = previewTemplate.innerHTML
        previewTemplate.parentNode.removeChild(previewTemplate);

        uploader = new Dropzone("#uploader", {
            url: staticSvelteEP.CHAT_IMG,
            maxFiles: 1,
            acceptedFiles: "image/*",
            thumbnailWidth: 60,
            thumbnailHeight: 60,
            previewTemplate: template,
            previewsContainer: gallery,
            clickable: ".clickable",
            renameFile(file:File):string {
                return fileName
            }
        });

        uploader.on("addedfile", (file:File) => {
            fileName = new Date().getTime() + '_' + file.name
        })

        uploader.on("removedfile", () => {
            fileName = ""
        })

        uploader.on("processing", () => {
            isDisabled = true
        })

        uploader.on("queuecomplete", () => {
            isDisabled = false
        })
    })


    function sendMsg() {
        message.trim()
        // if(message === "" || message.length === 0) {
        //     resetTextArea()
        //     return
        // }

        const comment = new ChatMessageModel<CommentModel>(
            chatMsgCat.COMMENT,
                new CommentModel(
                    message,
                    $currentUser.id,
                    $replyComment.createdAt != null ? $replyComment : null,
                    null,
                    true,
                    fileName
                )
        )
        console.log(comment)
        commentQueue.addComment(comment)
        resetTextArea()
    }

    function resetTextArea() {
        replyComment.close()
        message = ""
        textArea.style.height = "1.25rem"
        textArea.focus()
        fileName = ""
        uploader.removeAllFiles()
    }

    function sendOrResize(e:KeyboardEvent){
        textArea.style.height = "1px"
        textArea.style.height = `${textArea.scrollHeight}px`

        if(e.key == "Enter"){
            message = message.slice(0, -1)
            sendMsg()
        }
    }

    $: if($replyComment) {
        if($replyComment.text !== undefined && textArea !== undefined) {
            textArea.focus()
        }
    }
</script>

<div id="uploader">
    <div class="flex">
        <div class="flex flex-col-reverse">
            <button class="rounded-full py-3 clickable">
                <Image size="1.4em" class="clickable"/>
            </button>
        </div>
        <div class="flex-1
                    mx-3
                    border-solid
                    bg-canvas-secondary
                    border-2
                    border-gray-400
                    p-2
                    rounded-lg
                    shadow-sm">

            <ReplyComment />

            <div bind:this={gallery} class="dropzone-previews"></div>

            <textarea class="text-sm
                            h-5
                            p-0
                            overflow-hidden
                            border-0
                            focus:border-0
                            focus:outline-0"
                      name="msg"
                      bind:this={textArea}
                      bind:value={message}
                      on:keyup={e => sendOrResize(e)}></textarea>
        </div>

        <div class="flex flex-col-reverse">
            <button on:click={sendMsg} class="rounded-full py-3 {isDisabled}" disabled={isDisabled}>
                {#if isDisabled}
                    <Spinner size="sm" thickness="s" color="grey"/>
                {:else}
                    <Send size="1.4em"/>
                {/if}
            </button>
        </div>
    </div>
</div>


<!-- preview template -->
<div class="files" >
    <div id="template" class="file-row" bind:this={previewTemplate}>
        <!-- This is used as the file preview template -->
        <div class="flex">
            <div>
                <span class="preview relative">
                    <img data-dz-thumbnail />
                    <button data-dz-remove class="bg-transparent absolute top-2 left-[3.3rem]">
                        <CloseCircleOutline />
                    </button>
                </span>
            </div>
            <div class="flex flex-col-reverse data-dz-errormessage p-3"></div>
        </div>
    </div>
</div>
