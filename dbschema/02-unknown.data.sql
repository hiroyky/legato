use legato_db;

INSERT INTO genres
(`genre_id`, `name`, `name_hash`, `created_at`, `updated_at`)
VALUES
(
    1,
    "unknown",
    "ba8f0d3937ddaf252e41e89a1f9ae52b80a7e7347545098bdeab3d0aa90e865dc4056e7d69b3a623fb19beb2d9fb284089e688f99f6afa131b1bb4b053174246",
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

INSERT INTO album_artists
(`album_artist_id`, `name`, `name_hash`, `created_at`, `updated_at`)
VALUES
(
    1,
    "unknown",
    "ba8f0d3937ddaf252e41e89a1f9ae52b80a7e7347545098bdeab3d0aa90e865dc4056e7d69b3a623fb19beb2d9fb284089e688f99f6afa131b1bb4b053174246",
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

INSERT  INTO albums
(`album_id`, `name`, `album_artist_id`, `disc_no`, `disc_total`, `created_at`, `updated_at`)
VALUES
(
    1,
    "unknown",
    1,
    0,
    0,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);

