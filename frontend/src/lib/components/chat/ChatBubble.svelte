<script lang="ts">
    import {CommentModel} from "$lib/models/classes/comment.model";
    import {currentUser} from "$lib/stores/user.store";
    import {staticEP} from "$lib/models/enums/endpoints.enum";
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import {UserModel} from "$lib/models/classes/user.model";
    import ChatContent from "$lib/components/chat/ChatContent.svelte";

    export let comment:CommentModel = new CommentModel()
    export let user:UserModel = new UserModel()

    function isCurrentUser():boolean {
        return $currentUser.id === user.id
    }

    $: currentUserBubbleContainer = isCurrentUser() ? "ml-auto justify-end" : ""
    $: currentUserBubble = isCurrentUser() ? "bg-purple-800" : "bg-gray-300"
    $: roundedCorners = isCurrentUser() ? "rounded-l-lg rounded-br-lg" : "rounded-r-lg rounded-bl-lg"
</script>

{#if user.authLvl === authLvl.BOT}
    <div class="text-center mt-2 text-s">
        <p class="text-sm">{comment.text}</p>
    </div>
{:else}
    <div class="flex w-full mt-2 space-x-3 max-w-xs {currentUserBubbleContainer}">
        {#if !isCurrentUser()}
            <img class="flex-shrink-0 h-10 w-10 rounded-full" src={staticEP.IMG + user.icon} alt={user.name + "'s avatar"}>
        {/if}
        <div>

            <div class="p-3 {roundedCorners} {currentUserBubble}">
                <ChatContent comment={comment} user={user} />
            </div>

        </div>
        {#if isCurrentUser()}
            <img class="flex-shrink-0 h-10 w-10 rounded-full" src={staticEP.IMG + user.icon} alt={user.name + "'s avatar"}>
        {/if}
    </div>
{/if}