function BlogPost({ html, state }) {
  const { attrs, store } = state;
  const { key } = attrs;
  const { href, frontmatter } = store.posts[key];
  const { description = "", published = "", title = "" } = frontmatter;
  return html`
    <style>
      :host {
        display: block;
      }

      .avatar {
        width: 40px;
        aspect-ratio: 1 / 1;
      }
    </style>
    <a href="${href}" class="no-underline">
      <article class="pb2 mb-6">
        <div class="font-body leading3">
          <h1 class="font-heading font-bold leading1 text1 tracking-1 mbe0">
            ${title}
          </h1>
          <p class="mbe0">${description}</p>
          <p class="text-1 tracking1">${published}<br /></p>
        </div>
      </article>
    </a>
  `;
}
