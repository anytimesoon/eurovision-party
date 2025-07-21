<script lang="ts">
    import { currentUser } from "$lib/stores/user.store";

    interface Props {
        page: string;
    }

    let { page }: Props = $props();
</script>

{#if !$currentUser.isFriendOfFriend() }
    <ul class="p-3 flex justify-center">
        <li class="inline py-2 px-3 {page === 'settings' ? 'border-b-2 border-secondary' : ''}"><a href="/settings">Personal</a></li>
        {#if $currentUser.isAdmin()}
            <li class="inline py-2 px-3 {page === 'countries' ? 'border-b-2 border-secondary' : ''}"><a href="/settings/countries">Countries</a></li>
            <li class="inline py-2 px-3 {page === 'acts' ? 'border-b-2 border-secondary' : ''}"><a href="/settings/acts">Acts</a></li>
        {/if}
        <li class="inline py-2 px-3 {page === 'invites' ? 'border-b-2 border-secondary' : ''}"><a href="/settings/invites">Invites</a></li>
    </ul>
{/if}