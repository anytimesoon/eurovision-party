<script lang="ts">
    import type {CommentModel} from "$lib/models/classes/comment.model";
    import {staticSvelteEP} from "$lib/models/enums/endpoints.enum";
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import type {UserModel} from "$lib/models/classes/user.model";
    import ChatContent from "$lib/components/chat/ChatContent.svelte";
    import {swipeable} from '@react2svelte/swipeable';
    import type { SwipeEventData } from '@react2svelte/swipeable';
    import ReplyMenu from "$lib/components/chat/ReplyMenu.svelte";

    export let comment:CommentModel
    export let user:UserModel
    export let isCurrentUser:boolean
    export let openAvatarModal:Function = () => {}
    export let replyToComment:Function = () => {}
    let shouldShowReplyMenu:boolean = false

    function swipedHandler(e:CustomEvent<SwipeEventData>) {
        e.target.style.right = 0
        replyToComment(comment)
    }

    function swipingHandler(e:CustomEvent<SwipeEventData>) {
        if (e.detail.dir == "Right") {
            if (e.detail.deltaX < 100) {
                e.target.style.right = (e.detail.deltaX * -1) + "px"
            }
        }
    }

    const replyButtonHandler = () => {
        replyToComment(comment)
    }

    $: userNameStyle = isCurrentUser ? "text-right" : ""
    $: currentUserBubbleContainer = isCurrentUser ? "ml-auto justify-end" : ""
    $: currentUserImage = isCurrentUser ? "order-last ml-2" : "mr-2"
    $: currentUserBubble = isCurrentUser ? "bg-chat-bubble-me" : "bg-chat-bubble-you"
    $: roundedCorners = isCurrentUser ? "rounded-l-md rounded-r-sm" : "rounded-r-md rounded-l-sm"
    $: compactBubble = comment.isCompact ? "mt-[0.1rem]" : "mt-3"
    $: menuPadding = (shouldShowReplyMenu && comment.isCompact) ? "pt-5" : ""
</script>

{#if user.authLvl === authLvl.BOT}
    <div class="text-center mt-2 text-s p-3">
        <p class="text-sm">{comment.text}</p>
    </div>
{:else}

    <div use:swipeable
         on:swiping={swipingHandler}
         on:swipedright={swipedHandler}
         on:mouseenter={() => shouldShowReplyMenu = true}
         on:mouseleave={() => shouldShowReplyMenu = false}
         class="flex w-full max-w-[22rem] relative {currentUserBubbleContainer} {compactBubble}">

        <div class="flex-shrink">
            {#if !isCurrentUser}
                <div class="w-10 h-10 mr-1">
                    {#if !comment.isCompact}
                        <img on:click={() => openAvatarModal(user)}
                             class="flex-shrink-0 rounded-full cursor-pointer {currentUserImage}"
                             src={staticSvelteEP.IMG + user.icon} alt={user.name + "'s avatar"}>
                    {/if}
                </div>
            {/if}
        </div>


        <div class="max-w-[85%] {menuPadding}">
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

{/if}