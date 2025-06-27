<script lang="ts">
    import type {CommentModel} from "$lib/models/classes/comment.model";
    import {staticEP} from "$lib/models/enums/endpoints.enum";
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import type {UserModel} from "$lib/models/classes/user.model";
    import ChatContent from "$lib/components/chat/ChatContent.svelte";
    import {swipeable} from '@react2svelte/swipeable';
    import type { SwipeEventData } from '@react2svelte/swipeable';
    import ReplyMenu from "$lib/components/chat/ReplyMenu.svelte";
    import ImageLoader from "$lib/components/images/ImageLoader.svelte";
    import {replyComment} from "$lib/stores/replyComment.store";

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
    let shouldShowReplyMenu:boolean = $state(false)
    let bubble:HTMLDivElement = $state()

    function swipedHandler() {
        bubble.style.right = "0px"
        bubble.parentElement.classList.remove("overflow-y-hidden")
        bubble.parentElement.classList.add("overflow-y-auto")
    }

    function swipedRightHandler() {
        replyComment.set(comment)
    }

    function swipingHandler(e:CustomEvent<SwipeEventData>) {
        if (e.detail.dir == "Right") {
            if (e.detail.deltaX < 100) {
                bubble.parentElement.classList.add("overflow-y-hidden")
                bubble.parentElement.classList.remove("overflow-y-auto")
                bubble.style.right = (e.detail.deltaX * -1) + "px"
            }
        }
    }

    const replyButtonHandler = () => {
        replyComment.set(comment)
    }

    let userNameStyle = $derived(isCurrentUser ? "text-right" : "")
    let currentUserBubbleContainer = $derived(isCurrentUser ? "ml-auto justify-end" : "")
    let currentUserImage = $derived(isCurrentUser ? "order-last ml-2" : "mr-2")
    let currentUserBubble = $derived(isCurrentUser ? "bg-chat-bubble-me" : "bg-chat-bubble-you")
    let roundedCorners = $derived(isCurrentUser ? "rounded-l-md rounded-r-sm" : "rounded-r-md rounded-l-sm")
    let compactBubble = $derived(comment.isCompact ? "mt-[0.1rem]" : "mt-3")
    let menuPadding = $derived((shouldShowReplyMenu && comment.isCompact) ? "pt-5" : "")
</script>

{#if user && user.authLvl === authLvl.BOT}
    <div class="text-center mt-2 text-s p-3">
        <p class="text-sm">{comment.text}</p>
    </div>
{:else if user}

    <div use:swipeable
         onswiping={swipingHandler}
         onswipedright={swipedRightHandler}
         onswiped={swipedHandler}
         onmouseenter={() =>{shouldShowReplyMenu = true}}
         onmouseleave={() => shouldShowReplyMenu = false}
         bind:this={bubble}
         id="{comment.id}"
         class="flex w-full max-w-[22rem] relative {currentUserBubbleContainer} {compactBubble} transition-all"
         role=button
         tabindex="0">

        <div class="flex-shrink">
            {#if !isCurrentUser}
                <div class="w-10 h-10 mr-1 rounded-full overflow-hidden">
                    {#if !comment.isCompact}
                        <div onmouseup={() => openAvatarModal(user)} role="button" tabindex="0">
                            <ImageLoader customClasses="cursor-pointer {currentUserImage}"
                                         src={staticEP.AVATAR_IMG + user.icon} alt={user.name + "'s avatar"}/>
                        </div>
                    {/if}
                </div>
            {/if}
        </div>


        <div class="max-w-[75%] {menuPadding}">
            {#if !comment.isCompact && !isCurrentUser}
                <p class="block {userNameStyle}">{user.name}</p>
            {/if}

            <div class="px-3 py-2 {roundedCorners} {currentUserBubble} relative">
                {#if shouldShowReplyMenu}
                    <ReplyMenu replyButtonHandler={replyButtonHandler} />
                {/if}

                <ChatContent comment={comment}
                             isCurrentUser={isCurrentUser}/>
            </div>
        </div>
    </div>
    <div class="absolute">
        <div class="relative">

        </div>
    </div>
{:else}
<!--    <p>-->
<!--        If you're seeing this, it probably because your botId isn't set in the env vars.-->
<!--        If it keeps happening after it's set, try logging in again<br>-->
<!--        the user you are looking for is {user}<br><br>-->
<!--    </p>-->
{/if}