define({ "api": [
  {
    "type": "post",
    "url": "/api/v1/auth/login",
    "title": "login user",
    "version": "1.0.0",
    "name": "login",
    "group": "Auth",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>email address</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>password</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>message</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>user token</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "user",
            "description": "<p>user model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/auth.go",
    "groupTitle": "Auth"
  },
  {
    "type": "post",
    "url": "/api/v1/auth/register",
    "title": "register user",
    "version": "1.0.0",
    "name": "register",
    "group": "Auth",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>email address</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>password</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "repeat_password",
            "description": "<p>repeat password</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>message</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "token",
            "description": "<p>user token</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "user",
            "description": "<p>user model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/auth.go",
    "groupTitle": "Auth"
  },
  {
    "type": "post",
    "url": "/api/v1/categories",
    "title": "add category",
    "version": "1.0.0",
    "name": "addCategory",
    "group": "Category",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "parent",
            "description": "<p>parent</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "thumbnail",
            "description": "<p>thumbnail id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "category",
            "description": "<p>category model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/category.go",
    "groupTitle": "Category"
  },
  {
    "type": "put",
    "url": "/api/v1/categories/:id",
    "title": "edit category",
    "version": "1.0.0",
    "name": "editCategory",
    "group": "Category",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "parent",
            "description": "<p>parent</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "thumbnail",
            "description": "<p>thumbnail id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "category",
            "description": "<p>category model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/category.go",
    "groupTitle": "Category"
  },
  {
    "type": "get",
    "url": "/api/v1/categories/:id",
    "title": "get a category",
    "version": "1.0.0",
    "name": "getCategory",
    "group": "Category",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "category",
            "description": "<p>category model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/category.go",
    "groupTitle": "Category"
  },
  {
    "type": "get",
    "url": "/api/v1/categories",
    "title": "list of categories",
    "version": "1.0.0",
    "name": "listCategories",
    "group": "Category",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "q",
            "description": "<p>query search</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>list page</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "limit",
            "description": "<p>list limit</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "sort",
            "description": "<p>sort list example -created,title,...</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>page number</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "total_count",
            "description": "<p>total number of results</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "categories",
            "description": "<p>array of category model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/category.go",
    "groupTitle": "Category"
  },
  {
    "type": "delete",
    "url": "/api/v1/categories/:id",
    "title": "remove category",
    "version": "1.0.0",
    "name": "removeCategory",
    "group": "Category",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/category.go",
    "groupTitle": "Category"
  },
  {
    "type": "post",
    "url": "/api/v1/posts/:id/comments",
    "title": "add comment",
    "version": "1.0.0",
    "name": "addComment",
    "group": "Comment",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "content",
            "description": "<p>content</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "comment",
            "description": "<p>comment model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/comment.go",
    "groupTitle": "Comment"
  },
  {
    "type": "get",
    "url": "/api/v1/comments",
    "title": "list of comments",
    "version": "1.0.0",
    "name": "listComments",
    "group": "Comment",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "q",
            "description": "<p>query search</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "post",
            "description": "<p>post object id</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>list page</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "limit",
            "description": "<p>list limit</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "sort",
            "description": "<p>sort list example -created,title,...</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>page number</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "total_count",
            "description": "<p>total number of results</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "comments",
            "description": "<p>array of comment model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/comment.go",
    "groupTitle": "Comment"
  },
  {
    "type": "delete",
    "url": "/api/v1/commnets/:id",
    "title": "remove commnet",
    "version": "1.0.0",
    "name": "removeComment",
    "group": "Comment",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/comment.go",
    "groupTitle": "Comment"
  },
  {
    "type": "post",
    "url": "/api/v1/medias",
    "title": "add media",
    "version": "1.0.0",
    "name": "addMedia",
    "group": "Media",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "File",
            "optional": false,
            "field": "file",
            "description": "<p>media file</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "media",
            "description": "<p>media model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/media.go",
    "groupTitle": "Media"
  },
  {
    "type": "get",
    "url": "/api/v1/medias/:id",
    "title": "get a media",
    "version": "1.0.0",
    "name": "getMedia",
    "group": "Media",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "media",
            "description": "<p>media model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/media.go",
    "groupTitle": "Media"
  },
  {
    "type": "get",
    "url": "/api/v1/medias",
    "title": "list of medias",
    "version": "1.0.0",
    "name": "listMedias",
    "group": "Media",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>list page</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "limit",
            "description": "<p>list limit</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "sort",
            "description": "<p>sort list example -created,...</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>page number</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "total_count",
            "description": "<p>total number of results</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "medias",
            "description": "<p>array of media model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/media.go",
    "groupTitle": "Media"
  },
  {
    "type": "delete",
    "url": "/api/v1/medias/:id",
    "title": "remove media",
    "version": "1.0.0",
    "name": "removeMedia",
    "group": "Media",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/media.go",
    "groupTitle": "Media"
  },
  {
    "type": "post",
    "url": "/api/v1/posts",
    "title": "add post",
    "version": "1.0.0",
    "name": "addPost",
    "group": "Post",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "title",
            "description": "<p>post title</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "content",
            "description": "<p>post content</p>"
          },
          {
            "group": "Parameter",
            "type": "String[]",
            "optional": false,
            "field": "categories",
            "description": "<p>list of category id</p>"
          },
          {
            "group": "Parameter",
            "type": "String[]",
            "optional": false,
            "field": "tags",
            "description": "<p>list of tag</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "thumbnail",
            "description": "<p>thumbnail id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "post",
            "description": "<p>post model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/post.go",
    "groupTitle": "Post"
  },
  {
    "type": "put",
    "url": "/api/v1/posts/:id",
    "title": "edit post",
    "version": "1.0.0",
    "name": "editPost",
    "group": "Post",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "title",
            "description": "<p>post title</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "content",
            "description": "<p>post content</p>"
          },
          {
            "group": "Parameter",
            "type": "String[]",
            "optional": false,
            "field": "categories",
            "description": "<p>list of category id</p>"
          },
          {
            "group": "Parameter",
            "type": "String[]",
            "optional": false,
            "field": "tags",
            "description": "<p>list of tag</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "thumbnail",
            "description": "<p>thumbnail id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "post",
            "description": "<p>post model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/post.go",
    "groupTitle": "Post"
  },
  {
    "type": "get",
    "url": "/api/v1/posts/:id",
    "title": "get a post",
    "version": "1.0.0",
    "name": "getPost",
    "group": "Post",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "post",
            "description": "<p>post model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/post.go",
    "groupTitle": "Post"
  },
  {
    "type": "get",
    "url": "/api/v1/posts",
    "title": "list of posts",
    "version": "1.0.0",
    "name": "listPosts",
    "group": "Post",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "q",
            "description": "<p>query search</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "user",
            "description": "<p>user id</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "category",
            "description": "<p>category id</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>list page</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "limit",
            "description": "<p>list limit</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "sort",
            "description": "<p>sort list example -created,title,...</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>page number</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "total_count",
            "description": "<p>total number of results</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "posts",
            "description": "<p>array of post model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/post.go",
    "groupTitle": "Post"
  },
  {
    "type": "delete",
    "url": "/api/v1/posts/:id",
    "title": "remove post",
    "version": "1.0.0",
    "name": "removePost",
    "group": "Post",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/post.go",
    "groupTitle": "Post"
  },
  {
    "type": "get",
    "url": "/api/v1/settings",
    "title": "get setting",
    "version": "1.0.0",
    "name": "getSetting",
    "group": "Setting",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "settings",
            "description": "<p>settings model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/settings.go",
    "groupTitle": "Setting"
  },
  {
    "type": "put",
    "url": "/api/v1/settings",
    "title": "set setting",
    "version": "1.0.0",
    "name": "setSetting",
    "group": "Setting",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "settings",
            "description": "<p>setting model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/settings.go",
    "groupTitle": "Setting"
  },
  {
    "type": "post",
    "url": "/api/v1/users",
    "title": "add user",
    "version": "1.0.0",
    "name": "addUser",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "about",
            "description": "<p>about</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "email",
            "description": "<p>email</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>password</p>"
          },
          {
            "group": "Parameter",
            "type": "Boolean",
            "optional": false,
            "field": "super_user",
            "description": "<p>super user</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "thumbnail",
            "description": "<p>thumbnail</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "user",
            "description": "<p>user model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/user.go",
    "groupTitle": "User"
  },
  {
    "type": "patch",
    "url": "/api/v1/profile/change-password",
    "title": "change password for current user",
    "version": "1.0.0",
    "name": "changePassword",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "old_password",
            "description": "<p>old password</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "new_password",
            "description": "<p>new password</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "repeat_password",
            "description": "<p>repeat password</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/user.go",
    "groupTitle": "User"
  },
  {
    "type": "put",
    "url": "/api/v1/profile",
    "title": "edit current user",
    "version": "1.0.0",
    "name": "editProfile",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "about",
            "description": "<p>about</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "thumbnail",
            "description": "<p>thumbnail</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "user",
            "description": "<p>user model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/user.go",
    "groupTitle": "User"
  },
  {
    "type": "put",
    "url": "/api/v1/users/:id",
    "title": "edit user",
    "version": "1.0.0",
    "name": "editUser",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "about",
            "description": "<p>about</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "thumbnail",
            "description": "<p>thumbnail</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "user",
            "description": "<p>user model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/user.go",
    "groupTitle": "User"
  },
  {
    "type": "get",
    "url": "/api/v1/profile",
    "title": "get current user",
    "version": "1.0.0",
    "name": "getProfile",
    "group": "User",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "user",
            "description": "<p>user object</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/user.go",
    "groupTitle": "User"
  },
  {
    "type": "get",
    "url": "/api/v1/users/:id",
    "title": "get a user",
    "version": "1.0.0",
    "name": "getUser",
    "group": "User",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "user",
            "description": "<p>user model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/user.go",
    "groupTitle": "User"
  },
  {
    "type": "get",
    "url": "/api/v1/users",
    "title": "list of users",
    "version": "1.0.0",
    "name": "listUsers",
    "group": "User",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "q",
            "description": "<p>query search</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>list page</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "limit",
            "description": "<p>list limit</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "sort",
            "description": "<p>sort list example -created,title,...</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>page number</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "total_count",
            "description": "<p>total number of results</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "users",
            "description": "<p>array of user model</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/user.go",
    "groupTitle": "User"
  },
  {
    "type": "delete",
    "url": "/api/v1/users/:id",
    "title": "remove user",
    "version": "1.0.0",
    "name": "removeUser",
    "group": "User",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "message",
            "description": "<p>success message.</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "type": "String",
            "optional": false,
            "field": "error",
            "description": "<p>api error message</p>"
          }
        ]
      }
    },
    "filename": "server/v1/user.go",
    "groupTitle": "User"
  }
] });
