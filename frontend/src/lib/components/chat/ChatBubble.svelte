<script lang="ts">
    import type {CommentModel} from "$lib/models/classes/comment.model";
    import {staticEP} from "$lib/models/enums/endpoints.enum";
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import type {UserModel} from "$lib/models/classes/user.model";
    import ChatContent from "$lib/components/chat/ChatContent.svelte";
    import ReplyMenu from "$lib/components/chat/ReplyMenu.svelte";
    import ImageLoader from "$lib/components/images/ImageLoader.svelte";
    import {replyComment} from "$lib/stores/replyComment.store";
    import {emojiPickerState} from "$lib/stores/emojiPickerState.store";
    import VoteNotificationMessage from "$lib/components/chat/VoteNotificationMessage.svelte";
    import ReactionBelt from "$lib/components/chat/ReactionBelt.svelte";

    interface Props {
        comment: CommentModel;
        user: UserModel;
        isCurrentUser: boolean;
        openAvatarModal?: Function;
    }

    let {
        comment,
        user,
        isCurrentUser,
        openAvatarModal = () => {}
    }: Props = $props();
    let bubble:HTMLDivElement = $state()
    let touchStartX: number = $state(0)
    let currentDeltaX: number = $state(0)
    let rafId: number = $state(0)
    let isSwiping: boolean = $state(false)

    $effect(() => {
        if (!user) {
            console.error("If you're seeing this, it probably because your botId isn't set in the env vars. If it keeps happening after it's set, try logging in again. Comment: ", comment)
        }
    })

    const updateBubblePosition = () => {
        if (rafId !== null) {
            cancelAnimationFrame(rafId)
        }

        if (currentDeltaX < 0 && Math.abs(currentDeltaX) < 100) {
            bubble.style.right = `${currentDeltaX}px`;
        }

        if (isSwiping) { // Only request next frame if still swiping
            rafId = requestAnimationFrame(updateBubblePosition);
        } else {
            rafId = null;
        }

    }

    const handleTouchStart = (e:TouchEvent) => {
        touchStartX = e.touches[0].clientX
        isSwiping = true
        bubble.style.transition = "none"
        rafId = requestAnimationFrame(updateBubblePosition)
    }

    const handleTouchMove = (e:TouchEvent) => {
        if (!isSwiping) return

        const touchCurrentX = e.touches[0].clientX
        currentDeltaX = touchStartX - touchCurrentX
    }

    const handleTouchEnd = () => {
        if (!isSwiping) return

        if (rafId !== null) {
            cancelAnimationFrame(rafId)
            rafId = null
        }

        bubble.style.transition = "all 250ms cubic-bezier(0.4, 0, 0.2, 1)"
        if (currentDeltaX < -10) {
            replyComment.set(comment)
        }

        bubble.style.right = "0px"
        isSwiping = false
        currentDeltaX = 0
    }

    const replyButtonHandler = (e:MouseEvent) => {
        e.stopPropagation()
        e.preventDefault()
        replyComment.set(comment)
    }

    let currentUserBubbleContainer = $derived(isCurrentUser ? "ml-auto justify-end" : "")
    let currentUserImage = $derived(isCurrentUser ? "order-last ml-2" : "mr-2")
    let currentUserBubble = $derived(isCurrentUser ? "bg-chat-bubble-me" : "bg-chat-bubble-you")
    let roundedCorners = $derived(isCurrentUser ? "rounded-l-md rounded-r-sm" : "rounded-r-md rounded-l-sm")
    let compactBubble = $derived(comment.isCompact ? "mt-[0.1rem]" : "mt-7")
    let tooltipPosition = $derived(isCurrentUser ? "right-[4.5rem]" : "left-[6.5rem]")
</script>

    {#if user && user.authLvl === authLvl.BOT && comment.isVoteNotification}
        <VoteNotificationMessage />
    {:else if user && user.authLvl === authLvl.BOT}
        <div class="text-center mt-2 text-s p-3">
            <p class="text-sm">{comment.text}</p>
        </div>
    {:else if user}
        <div
             ontouchstart={handleTouchStart}
             ontouchmove={handleTouchMove}
             ontouchend={handleTouchEnd}
             ontouchcancel={handleTouchEnd}
             oncontextmenu={(e:MouseEvent) => {
                 e.preventDefault()
                 emojiPickerState.open(comment)
             }}
             bind:this={bubble}
             id="{comment.id}"
             class="flex w-full max-w-[22rem] {currentUserBubbleContainer} {compactBubble} relative group break-all"
             role=button
             tabindex="0">

            <span class="absolute z-20 opacity-0 transition-opacity duration-300 group-hover:opacity-100 {tooltipPosition} -top-6">
                <ReplyMenu replyButtonHandler={replyButtonHandler} parentComment={comment}/>
            </span>

            <div class="flex-shrink z-10">
                {#if !isCurrentUser}
                    <div class="w-8 h-8">
                        {#if !comment.isCompact}
                            <div>
                                <div onmouseup={() => openAvatarModal(user)} role="button" tabindex="0" class="border-4 border-canvas-secondary overflow-hidden rounded-full relative -top-5">
                                    <ImageLoader customClasses="cursor-pointer {currentUserImage}"
                                                 src={staticEP.AVATAR_IMG + user.icon} alt={user.name + "'s avatar"}/>
                                </div>
                                <span class="text-xs absolute -top-4 left-8">{user.name}</span>
                            </div>
                        {/if}
                    </div>
                {/if}
            </div>

            <div class="px-2 py-1.5 {roundedCorners} {currentUserBubble} relative -ml-3 max-w-[85%]">
                <ChatContent comment={comment}
                             isCurrentUser={isCurrentUser}/>
            </div>
        </div>

        {#if comment.reactions.size > 0}
            <ReactionBelt comment={comment} isCurrentUser={isCurrentUser}/>
        {/if}
    {:else}
<!--No operation. See the error message in the console-->
    {/if}