#include <stdio.h>           // Needed To print stuff to the console
#include <json-c/json.h>     // Needed To parse the json
#include <hiredis/hiredis.h> // Needed to connect to redis

// The json string to benchmark
const char* bench = "{\"user\":356091260429402122,\"name\":\"Why are you reading\",\"money\":9164,\"xp\":6000000,\"pvpwins\":14,\"money_booster\":0,\"time_booster\":0,\"luck_booster\":0,\"marriage\":463318425901596672,\"background\":\"https://i.imgur.com/LRV2QCK.png\",\"guild\":15306,\"class\":[\"Paragon\",\"White Sorcerer\"],\"deaths\":0,\"completed\":0,\"lovescore\":647,\"guildrank\":\"Leader\",\"backgrounds\":null,\"puzzles\":0,\"atkmultiply\":\"10.0\",\"defmultiply\":\"10.0\",\"crates_common\":30,\"crates_uncommon\":2,\"crates_rare\":1,\"crates_magic\":0,\"crates_legendary\":0,\"luck\":\"1.0\",\"god\":null,\"favor\":0,\"race\":\"Elf\",\"cv\":2,\"reset_points\":2,\"chocolates\":0,\"trickortreat\":0,\"eastereggs\":0,\"colour\":{\"red\":255,\"green\":255,\"blue\":255,\"alpha\":0.8}}";

// Main Function
int main( int argc, char** argv )
{
    char *newjson;
    newjson = malloc(sizeof(bench) + 32);                                       // allocate the memory for the jobj string
    
    redisReply   *reply;                                                  // Redis reply
    redisContext *context = redisConnect("LOCALHOST", 6379);              // Connect to redis and start a redis context
    if (!context) { printf("Failed to connect to redis\n"); return -1; }  // Check to make sure the connection was successfull

    // SET the key "bench" to the benchmark string
    reply = redisCommand(context, "SET %s %s", "bench", bench);
    if(!reply || context->err) { 
        if (context->err)                                        // If there was an error message
        { printf("Failed with err %s\n", context->errstr); redisFree(context); return -1; }
        else                                                     // else (no connection probably)
        {  printf("Can't send command to Redis\n"); return -1; }
    }
    freeReplyObject(reply);                                      // Free the reply object

    for (int i = 0; i < 100000; i++) {
        reply = redisCommand(context, "GET bench");                             // Get the current json from redis
        if (!reply || context->err || reply->type !=REDIS_REPLY_STRING)         // Check if we actually got it
        { printf("ERROR \n----\n%s\n in iter %d\n", context->errstr, i); }      // Print to console in the case that we didnt get it

        json_object *jobj = json_tokener_parse(reply->str);                     // Parse the json string
        json_object *crateCommon;   json_object_object_get_ex(jobj, "crates_common", &crateCommon);     // Get the common crate json object
        json_object *crateUncommon; json_object_object_get_ex(jobj, "crates_uncommon", &crateUncommon); // Get the uncommon crate json object
        int common, uncommon;
        common = json_object_get_int(crateCommon); uncommon = json_object_get_int(crateUncommon);       // Grab the current number of common and uncommon crates
        json_object_object_del(jobj, "crates_common"); json_object_object_del(jobj, "crates_uncommon"); // Delete the common and uncommon crate fields
        json_object *jCommonInt = json_object_new_int(common + 1);                                      // Create new json int object for common
        json_object *jUncommonInt = json_object_new_int(uncommon + 1);                                  // Create new json int object for uncommon
        json_object_object_add(jobj, "crates_common", jCommonInt);                                      // Add the common int json object field to the jobj
        json_object_object_add(jobj, "crates_uncommon", jUncommonInt);                                  // Add the uncommon int json object field to the jobj
        newjson = json_object_get_string(jobj);                                                         // Convert the jobj to string

        reply = redisCommand(context, "SET bench %s", newjson);                      // Send the redis command to set the new json string
        if (!reply || context->err)                                                  // Check for an error
        {
            if(context->err)                                                         // If have error message
            { printf("REDIS ERROR on iter %d\n Error: %s\n", i, context->errstr); }  // Print the error message to the console
            else                                                                     // else (meaning no connection to redis)
            { printf("FAILED to send redis command on iter %d\n", i); }              // Print error message to console
        }

        // free(newjson);          // Free the memory for the string
        freeReplyObject(reply); // Free the reply object
    }
    
    free(newjson);
    redisFree(context);                     // Disconnect and free the redis context
    return 0;                               // Exit the program
}
