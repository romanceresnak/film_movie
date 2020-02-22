#Film movie project #

## Install dependencies ##    
   * go get -u google.golang.org/grpc    
   * go get -u github.com/golang/protobuf/protoc-gen-go
   
## Database tables ##
CREATE TABLE actor (    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;actor_id    SMALLINT     PK   
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;first_name  VARCHAR(45)  NOT NULL,    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;last_name   VARCHAR(45)  NOT NULL,    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;last_update TIMESTAMP    NOT NULL     
) 

CREATE TABLE language (     
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;language_id  TINYINT    PK,    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;name         CHAR(20)   NOT NULL,     
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;last_update  TIMESTAMP  NOT NULL      
)

CREATE TABLE film (     
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;film_id              SMALLINT     pk,  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;title                VARCHAR(255) NOT NULL,   
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;description          TEXT         DEFAULT NULL,     
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;release_year         YEAR         DEFAULT NULL,       
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;language_id          TINYINT      UNSIGNED NOT NULL,  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;original_language_id TINYINT      UNSIGNED DEFAULT NULL,
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;rental_duration      TINYINT      UNSIGNED NOT NULL DEFAULT 3,
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;rental_rate          DECIMAL(4,2) NOT NULL DEFAULT 4.99,  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;length               SMALLINT     UNSIGNED DEFAULT NULL,    
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;replacement_cost     DECIMAL(5,2) NOT NULL,  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;rating               Varchar(255),   
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;last_update          TIMESTAMP    NOT NULL  
) 