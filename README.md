# avito-test-trainee

### База данных
```mermaid

erDiagram
    
Banners ||--|| Features : ""
Banners ||--|{ BannerTags : ""
Tags ||--|{ BannerTags : ""

Banners {
        int id
        data jsonb
        bool is_enabled
        int feature_id
        string created_at
        string updated_at
    }

    Tags {
        int id
        string title
        string created_at
        string updated_at
    }

    Features {
        int id
        string title
        string created_at
        string updated_at
    }

    BannerTags {
        int banner_id
        int tag_id
    }
```