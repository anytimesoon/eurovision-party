<script lang="ts">
    export let page:string


</script>

<ul class="mx-auto max-w-max p-3 flex justify-between">
    <li class="inline py-2 px-3 {page === 'settings' ? 'border-b-2 border-secondary' : ''}"><a href="/settings">Personal</a></li>
    <li class="inline py-2 px-3"><a href="/admin/countries">Countries</a></li>
    <li class="inline py-2 px-3"><a href="/admin/acts">Acts</a></li>
    <li class="inline py-2 px-3"><a href="/admin/friends">Friends</a></li>
</ul>