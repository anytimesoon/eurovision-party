<script lang="ts">
    import type {CommentModel} from "$lib/models/classes/comment.model";
    import Spinner from "$lib/components/Spinner.svelte";
    import {socketRetryCount} from "$lib/stores/socketRetryCount.store";
    import {commentQueue} from "$lib/stores/commentQueue.store";
    import {userStore} from "$lib/stores/user.store";
    import CloseCircleOutline from "svelte-material-icons/CloseCircleOutline.svelte";
    import ImageLoader from "$lib/components/images/ImageLoader.svelte";
    import {staticEP} from "$lib/models/enums/endpoints.enum.js";
    import {emojiRegex, urlRegex} from "$lib/utils/regex";
    import VideoLoader from "$lib/components/images/VideoLoader.svelte";

    interface Props {
        comment: CommentModel;
        isCurrentUser: boolean;
    }

    let { comment, isCurrentUser }: Props = $props();
    let commentElement:HTMLParagraphElement = $state(document.createElement("p"))
    const SUPPORTED_VIDEO_TYPES = ['mp4', 'webm']

    $effect(() => {
        if(comment) {
            const linkifiedText = comment.text.replace(urlRegex, function(url:string) {
                return '<a href="' + url + '" target="_blank" rel="noopener noreferrer">' + url + '</a>'
            })

            const emojiCount = linkifiedText.match(emojiRegex) ? linkifiedText.match(emojiRegex).length : 0
            const noEmojis = linkifiedText.replace(emojiRegex, "").length

            const processedText = linkifiedText.replace(emojiRegex, function(emoji:string) {
                if(emojiCount <= 3 && noEmojis === 0) {
                    return '<span class="text-[3em] -tracking-[0.35em]">' + emoji + '</span>'
                }
                return '<span class="text-lg -tracking-[0.25em]">' + emoji + '</span>'
            })
            const parser = new DOMParser
            const htmlDoc = parser.parseFromString(processedText, "text/html")
            commentElement.innerHTML = htmlDoc.body.innerHTML
        }
    });

    function removeMessage(e: Event) {
        e.preventDefault()
        commentQueue.removeMessage(comment.id)
    }

    let contentTextStyle = $derived(isCurrentUser ? "text-typography-chat-me" : "text-typography-chat-you")
</script>

<div class="flex">
    <div>
        {#if comment.replyToComment}
            <a href="#{comment.replyToComment.id}">
                <div class="text-sm bg-canvas-primary rounded px-3 py-1 mb-1">
                    <p class="text-xs">{$userStore.get(comment.replyToComment.userId).name}</p>
                    {#if comment.replyToComment.fileName !== ""}
                        <div class="h-[40px] rounded overflow-hidden">
                            {#if SUPPORTED_VIDEO_TYPES.includes(comment.fileName)}
                                <VideoLoader src={staticEP.CHAT_IMG+comment.replyToComment.fileName} customClasses=""/>
                            {:else}
                                <ImageLoader src={staticEP.CHAT_IMG+comment.replyToComment.fileName} alt="comment image" customClasses="h-[40px]"/>
                            {/if}
                        </div>
                    {/if}
                    <span class="text-typography-chat-you pt-1 block">
                        {comment.replyToComment.text}
                    </span>
                </div>
            </a>
        {/if}

        {#if comment.fileName !== ""}
            <div class="mb-1 rounded overflow-hidden">
                {#if SUPPORTED_VIDEO_TYPES.includes(comment.fileName)}
                    <VideoLoader src={staticEP.CHAT_IMG + comment.fileName} customClasses=""/>
                {:else}
                    <ImageLoader src={staticEP.CHAT_IMG + comment.fileName} alt="comment image" customClasses=""/>
                {/if}
            </div>
        {/if}

        <p bind:this={commentElement} class="{contentTextStyle} pr-8"></p>
    </div>


    <div class="flex flex-col-reverse flex-shrink">
        {#if comment.createdAt !== null}
            <span class="text-[0.6rem] block {contentTextStyle} text-right break-normal">
                {
                    comment.createdAt.getHours() + ":" +
                    (comment.createdAt.getMinutes() < 10 ?
                            "0" + comment.createdAt.getMinutes() :
                            comment.createdAt.getMinutes()
                    )
                }
            </span>
        {:else }
            <span class="text-[0.6rem] block {contentTextStyle} text-right">
                <Spinner size={"s"} thickness={"s"} color={"grey"}/>
            </span>

            {#if $socketRetryCount > 2}
                <button onclick={e => removeMessage(e)}>
                    <span class="block p-0 m-0 -mt-2.5 text-warning cursor-pointer">
                        <CloseCircleOutline />
                    </span>
                </button>
            {/if}
        {/if}


    </div>
</div>

